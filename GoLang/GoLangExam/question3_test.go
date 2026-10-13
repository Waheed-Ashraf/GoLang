package main

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

type taskHandlerFunc func(context.Context, Task) error

func (f taskHandlerFunc) Handle(ctx context.Context, task Task) error { return f(ctx, task) }

func TestTaskRetriesAndWrapping(t *testing.T) {
	rootErr := errors.New("temporary service failure")
	for _, tc := range []struct {
		name      string
		failures  int
		attempts  int
		wantError bool
	}{
		{name: "eventual success", failures: 2, attempts: 3},
		{name: "retry limit", failures: 10, attempts: 4, wantError: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			calls := 0
			handler := taskHandlerFunc(func(context.Context, Task) error {
				calls++
				if calls <= tc.failures {
					return fmt.Errorf("send: %w", &RetryableError{Err: rootErr})
				}
				return nil
			})
			result := processTask(context.Background(), Task{ID: 1, Name: "email"}, handler)
			if calls != tc.attempts || result.Attempts != tc.attempts || result.Retries != tc.attempts-1 {
				t.Fatalf("unexpected attempt counts: calls=%d result=%+v", calls, result)
			}
			if tc.wantError {
				var retryErr *RetryableError
				if !errors.Is(result.Err, rootErr) || !errors.As(result.Err, &retryErr) {
					t.Fatalf("lost wrapped error: %v", result.Err)
				}
			} else if result.Err != nil {
				t.Fatalf("successful retry retained an error: %v", result.Err)
			}
		})
	}
}

func TestPanicDoesNotKillWorker(t *testing.T) {
	tasks := make(chan Task, 3)
	for id := 1; id <= 3; id++ {
		tasks <- Task{ID: id, Name: "test"}
	}
	close(tasks)
	handler := taskHandlerFunc(func(_ context.Context, task Task) error {
		if task.ID == 1 {
			panic("bad task")
		}
		if task.ID == 2 {
			return errors.New("permanent error")
		}
		return nil
	})
	results, stats, err := RunTaskPool(context.Background(), tasks, 1, map[string]Handler{"test": handler})
	if err != nil || len(results) != 3 || stats.Succeeded != 1 || stats.Failed != 2 || stats.Panicked != 1 || stats.Retried != 0 {
		t.Fatalf("unexpected pool output: results=%+v stats=%+v err=%v", results, stats, err)
	}
	var panicErr *PanicError
	if !errors.As(results[0].Err, &panicErr) || results[2].Err != nil {
		t.Fatal("panic was not recovered or the worker did not continue")
	}
}

func TestPoolConcurrentStats(t *testing.T) {
	tasks := make(chan Task, 90)
	for id := 0; id < 90; id++ {
		tasks <- Task{ID: id, Name: "test"}
	}
	close(tasks)
	handler := taskHandlerFunc(func(_ context.Context, task Task) error {
		switch task.ID % 3 {
		case 1:
			return errors.New("permanent failure")
		case 2:
			panic("handler panic")
		}
		return nil
	})
	results, stats, err := RunTaskPool(context.Background(), tasks, 8, map[string]Handler{"test": handler})
	if err != nil || len(results) != 90 || stats.Succeeded != 30 || stats.Failed != 60 || stats.Panicked != 30 {
		t.Fatalf("incorrect concurrent counts: %+v, results=%d, err=%v", stats, len(results), err)
	}
	seen := make(map[int]bool)
	var total time.Duration
	for _, result := range results {
		if seen[result.Task.ID] {
			t.Fatalf("duplicate result for task %d", result.Task.ID)
		}
		seen[result.Task.ID] = true
		total += result.Duration
	}
	if stats.TotalDuration != total || stats.AverageDuration() != total/90 {
		t.Fatal("incorrect duration aggregation")
	}
}

func TestPoolCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	tasks := make(chan Task, 1) // Intentionally left open to test shutdown.
	tasks <- Task{ID: 1, Name: "test"}
	started := make(chan struct{})
	handler := taskHandlerFunc(func(ctx context.Context, _ Task) error {
		close(started)
		<-ctx.Done()
		return ctx.Err()
	})
	done := make(chan struct{})
	var results []TaskResult
	var poolErr error
	go func() {
		results, _, poolErr = RunTaskPool(ctx, tasks, 2, map[string]Handler{"test": handler})
		close(done)
	}()
	select {
	case <-started:
	case <-time.After(time.Second):
		t.Fatal("handler did not start")
	}
	cancel()
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("workers did not shut down")
	}
	if !errors.Is(poolErr, context.Canceled) || len(results) != 1 || !errors.Is(results[0].Err, context.Canceled) {
		t.Fatalf("cancellation was lost: poolErr=%v results=%+v", poolErr, results)
	}
}

func TestDemoHandlersAndInvalidInput(t *testing.T) {
	tasks := make(chan Task, 6)
	tasks <- Task{ID: 1, Name: "email", Payload: map[string]any{"to": "test@example.com"}}
	tasks <- Task{ID: 2, Name: "resize", Payload: map[string]any{"path": "photo.jpg", "width": 100}}
	tasks <- Task{ID: 3, Name: "report", Payload: map[string]any{"title": "Report"}}
	tasks <- Task{ID: 4, Name: "unknown"}
	tasks <- Task{ID: 5, Name: "resize", Payload: map[string]any{"path": "photo.jpg", "width": -1}}
	tasks <- Task{ID: 6, Name: "email", Payload: map[string]any{"to": "test@example.com", "temporary_failure": true}}
	close(tasks)
	handlers := map[string]Handler{"email": EmailHandler{}, "resize": ResizeImageHandler{}, "report": ReportHandler{}}
	results, stats, err := RunTaskPool(context.Background(), tasks, 3, handlers)
	if err != nil || len(results) != 6 || stats.Succeeded != 3 || stats.Failed != 3 || stats.Retried != 1 || stats.RetryAttempts != 3 {
		t.Fatalf("unexpected demo results: %+v, err=%v", stats, err)
	}
	if _, _, err := RunTaskPool(context.Background(), nil, 0, handlers); err == nil {
		t.Fatal("zero workers accepted")
	}
	if (TaskStats{}).AverageDuration() != 0 {
		t.Fatal("empty stats must have zero average")
	}
}
