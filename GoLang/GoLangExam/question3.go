package main

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

// Name identifies the task type in the handler registry.
// Payload must not be modified after a task is submitted.
type Task struct {
	ID       int
	Name     string
	Payload  map[string]any
	Priority int
}

type Handler interface {
	Handle(ctx context.Context, t Task) error
}

type RetryableError struct{ Err error }

func (e *RetryableError) Error() string { return "temporary failure: " + e.Err.Error() }
func (e *RetryableError) Unwrap() error { return e.Err }

type ValidationError struct{ Message string }

func (e *ValidationError) Error() string { return e.Message }

type PanicError struct {
	TaskID int
	Value  any
}

func (e *PanicError) Error() string {
	return fmt.Sprintf("task %d panicked: %v", e.TaskID, e.Value)
}

type EmailHandler struct{}
type ResizeImageHandler struct{}
type ReportHandler struct{}

func (EmailHandler) Handle(ctx context.Context, t Task) error {
	if _, err := requiredString(t, "to"); err != nil {
		return fmt.Errorf("email: %w", err)
	}
	return simulateWork(ctx, t)
}

func (ResizeImageHandler) Handle(ctx context.Context, t Task) error {
	if _, err := requiredString(t, "path"); err != nil {
		return fmt.Errorf("resize image: %w", err)
	}
	width, ok := t.Payload["width"].(int)
	if !ok || width <= 0 {
		return &ValidationError{Message: "resize image: width must be a positive int"}
	}
	return simulateWork(ctx, t)
}

func (ReportHandler) Handle(ctx context.Context, t Task) error {
	if _, err := requiredString(t, "title"); err != nil {
		return fmt.Errorf("report: %w", err)
	}
	return simulateWork(ctx, t)
}

func requiredString(t Task, key string) (string, error) {
	value, ok := t.Payload[key].(string)
	if !ok || value == "" {
		return "", &ValidationError{Message: key + " must be a nonempty string"}
	}
	return value, nil
}

// These handlers simulate operations so the exam example needs no external services.
// Payload flags make errors and panics reproducible.
func simulateWork(ctx context.Context, t Task) error {
	timer := time.NewTimer(15 * time.Millisecond)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
	}
	if shouldPanic, _ := t.Payload["panic"].(bool); shouldPanic {
		panic("unexpected handler failure")
	}
	if temporary, _ := t.Payload["temporary_failure"].(bool); temporary {
		return &RetryableError{Err: errors.New("service unavailable")}
	}
	if permanent, _ := t.Payload["permanent_failure"].(bool); permanent {
		return errors.New("operation rejected")
	}
	return nil
}

type TaskResult struct {
	Task     Task
	Attempts int
	Retries  int
	Panicked bool
	Duration time.Duration // Includes all attempts and retry waits.
	Err      error
}

type TaskStats struct {
	Succeeded     int
	Failed        int
	Retried       int // Number of tasks that were retried at least once.
	RetryAttempts int // Total additional attempts across all tasks.
	Panicked      int // Panicked tasks are also counted as failed.
	TotalDuration time.Duration
}

func (s TaskStats) AverageDuration() time.Duration {
	count := s.Succeeded + s.Failed
	if count == 0 {
		return 0
	}
	return s.TotalDuration / time.Duration(count)
}

type taskCollector struct {
	mu      sync.Mutex
	results []TaskResult
	stats   TaskStats
}

func (c *taskCollector) add(r TaskResult) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.results = append(c.results, r)
	if r.Err == nil {
		c.stats.Succeeded++
	} else {
		c.stats.Failed++
	}
	if r.Retries > 0 {
		c.stats.Retried++
	}
	c.stats.RetryAttempts += r.Retries
	if r.Panicked {
		c.stats.Panicked++
	}
	c.stats.TotalDuration += r.Duration
}

// RunTaskPool waits until all workers exit before returning the collected results.
// Close tasks after sending the final task. Cancellation stops taking new tasks;
// in-flight handlers and retry waits receive the same canceled context.
// Custom handlers must observe ctx for prompt shutdown. Keep the registry immutable
// while the pool runs, and make shared handler implementations concurrency-safe.
func RunTaskPool(ctx context.Context, tasks <-chan Task, workers int, handlers map[string]Handler) ([]TaskResult, TaskStats, error) {
	if workers <= 0 {
		return nil, TaskStats{}, &ValidationError{Message: "workers must be positive"}
	}
	collector := taskCollector{results: make([]TaskResult, 0)}
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for {
				if ctx.Err() != nil {
					return
				}
				select {
				case <-ctx.Done():
					return
				case task, ok := <-tasks:
					if !ok || ctx.Err() != nil {
						return
					}
					collector.add(processTask(ctx, task, handlers[task.Name]))
				}
			}
		}()
	}
	wg.Wait() // All collector writes finish before these reads.
	return collector.results, collector.stats, ctx.Err()
}

const maxTaskRetries = 3 // One initial attempt plus at most three retries.

func processTask(ctx context.Context, t Task, handler Handler) (result TaskResult) {
	result.Task = t
	start := time.Now()
	defer func() { result.Duration = time.Since(start) }()
	if handler == nil {
		result.Err = &ValidationError{Message: "unknown task type: " + t.Name}
		return
	}
	for {
		if err := ctx.Err(); err != nil {
			result.Err = fmt.Errorf("task %d canceled: %w", t.ID, err)
			return
		}
		result.Attempts++
		result.Retries = result.Attempts - 1
		err := invokeHandler(ctx, handler, t)
		if err == nil {
			return
		}
		result.Err = fmt.Errorf("task %d (%s), attempt %d: %w", t.ID, t.Name, result.Attempts, err)
		var panicErr *PanicError
		if errors.As(err, &panicErr) {
			result.Panicked = true
			return
		}
		var retryErr *RetryableError
		if !errors.As(err, &retryErr) || result.Retries == maxTaskRetries || ctx.Err() != nil {
			return
		}
		// Cancellation interrupts the backoff instead of waiting for a sleep.
		timer := time.NewTimer(time.Duration(result.Attempts) * 20 * time.Millisecond)
		select {
		case <-ctx.Done():
			timer.Stop()
			result.Err = fmt.Errorf("task %d canceled during retry: %w", t.ID, ctx.Err())
			return
		case <-timer.C:
		}
		result.Err = nil
	}
}

// Recover inside a helper called for each attempt, in the worker's goroutine.
// Returning from this helper lets the worker continue to its next task.
func invokeHandler(ctx context.Context, handler Handler, t Task) (err error) {
	defer func() {
		if value := recover(); value != nil {
			err = &PanicError{TaskID: t.ID, Value: value}
		}
	}()
	return handler.Handle(ctx, t)
}

func runQuestion3Demo(ctx context.Context, workers int) error {
	handlers := map[string]Handler{
		"email":  EmailHandler{},
		"resize": ResizeImageHandler{},
		"report": ReportHandler{},
	}
	batch := []Task{
		{ID: 1, Name: "email", Payload: map[string]any{"to": "student@example.com"}, Priority: 3},
		{ID: 2, Name: "resize", Payload: map[string]any{"path": "photo.jpg", "width": 800}, Priority: 2},
		{ID: 3, Name: "report", Payload: map[string]any{"title": "Exam results"}, Priority: 1},
		{ID: 4, Name: "email", Payload: map[string]any{"to": "student@example.com", "temporary_failure": true}, Priority: 3},
		{ID: 5, Name: "resize", Payload: map[string]any{"path": "photo.jpg", "width": -1}, Priority: 2},
		{ID: 6, Name: "report", Payload: map[string]any{"title": "Broken report", "panic": true}, Priority: 1},
		{ID: 7, Name: "email", Payload: map[string]any{"to": "student@example.com", "permanent_failure": true}, Priority: 1},
		{ID: 8, Name: "report", Payload: map[string]any{"title": "Report after panic"}, Priority: 0},
	}
	// Submit higher priorities first; concurrent completion order can vary.
	sort.SliceStable(batch, func(i, j int) bool { return batch[i].Priority > batch[j].Priority })
	tasks := make(chan Task, len(batch))
	for _, task := range batch {
		tasks <- task
	}
	close(tasks)
	results, stats, err := RunTaskPool(ctx, tasks, workers, handlers)
	sort.Slice(results, func(i, j int) bool { return results[i].Task.ID < results[j].Task.ID })
	for _, r := range results {
		status := "succeeded"
		if r.Err != nil {
			status = "failed"
		}
		fmt.Printf("Task %d (%s): %s, attempts=%d, retries=%d, panicked=%t, duration=%s\n",
			r.Task.ID, r.Task.Name, status, r.Attempts, r.Retries, r.Panicked, r.Duration.Round(time.Millisecond))
		if r.Err != nil {
			fmt.Printf("  %v\n", r.Err)
		}
	}
	fmt.Printf("\nSummary: succeeded=%d failed=%d retried=%d panicked=%d retry_attempts=%d average_duration=%s\n",
		stats.Succeeded, stats.Failed, stats.Retried, stats.Panicked, stats.RetryAttempts, stats.AverageDuration().Round(time.Millisecond))
	return err
}
