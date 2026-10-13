package main

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

// 1. Keep indices in a deque whose values decrease from front to back.
// Each index enters and leaves once, so this takes O(n) time.
func SlidingWindowMax(nums []int, k int) []int {
	if k <= 0 || k > len(nums) {
		return []int{}
	}
	deque := make([]int, 0, len(nums))
	result := make([]int, 0, len(nums)-k+1)
	for i, value := range nums {
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= value {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}
	return result
}

// 2. Anagrams share the same sorted letters. Preserve input/group order.
func GroupAnagrams(words []string) [][]string {
	groups := make([][]string, 0)
	indices := make(map[string]int)
	for _, word := range words {
		letters := []rune(word)
		sort.Slice(letters, func(i, j int) bool { return letters[i] < letters[j] })
		key := string(letters)
		index, found := indices[key]
		if !found {
			index = len(groups)
			indices[key] = index
			groups = append(groups, []string{})
		}
		groups[index] = append(groups[index], word)
	}
	return groups
}

// 3. Students and their average grades.
type Student struct {
	Name   string
	ID     int
	Grades map[string]float64
}

func (s Student) Average() float64 {
	if len(s.Grades) == 0 {
		return 0
	}
	var total float64
	for _, grade := range s.Grades {
		total += grade
	}
	return total / float64(len(s.Grades))
}

func TopStudent(students []Student) *Student {
	if len(students) == 0 {
		return nil
	}
	best := 0
	bestAverage := students[0].Average()
	for i := 1; i < len(students); i++ {
		average := students[i].Average()
		if average > bestAverage {
			best, bestAverage = i, average
		}
	}
	return &students[best] // Pointer to the original slice element.
}

// 4. sort.Interface requires Len, Less, and Swap.
type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (p ByAge) Len() int           { return len(p) }
func (p ByAge) Less(i, j int) bool { return p[i].Age < p[j].Age }
func (p ByAge) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Ranker is our custom contract: smaller ranks come first.
type Ranker interface {
	Len() int
	Rank(i int) int
	Swap(i, j int)
}

type rankAdapter struct{ Ranker }

func (r rankAdapter) Less(i, j int) bool { return r.Rank(i) < r.Rank(j) }

func SortRanked(r Ranker) { sort.Sort(rankAdapter{r}) }

func (p ByAge) Rank(i int) int { return p[i].Age }

// 5. Retry with waits of 100ms, 200ms, 400ms, ... capped at 2 seconds.
// There is no wait after success or after the final failure.
func Retry(times int, fn func() error) error {
	if times <= 0 {
		return errors.New("times must be positive")
	}
	if fn == nil {
		return errors.New("retry function must not be nil")
	}
	delay := 100 * time.Millisecond
	var lastErr error
	for attempt := 0; attempt < times; attempt++ {
		lastErr = fn()
		if lastErr == nil {
			return nil
		}
		if attempt < times-1 {
			time.Sleep(delay)
			delay *= 2
			if delay > 2*time.Second {
				delay = 2 * time.Second
			}
		}
	}
	return fmt.Errorf("after %d attempts: %w", times, lastErr)
}

// 6. An unrecovered panic in ANY goroutine crashes the whole program.
// Recovery must run in a deferred function in that same goroutine.
// Recover per job so the worker can continue processing subsequent jobs.
func runJob(job int, handle func(int) error) (err error) {
	defer func() {
		if recovered := recover(); recovered != nil {
			err = fmt.Errorf("job %d panicked: %v", job, recovered)
		}
	}()
	if err = handle(job); err != nil {
		return fmt.Errorf("job %d: %w", job, err)
	}
	return nil
}

// Callers must drain the returned error channel while workers run.
func WorkerPool(jobs <-chan int, workers int, handle func(int) error) <-chan error {
	errorsCh := make(chan error)
	if workers <= 0 || handle == nil {
		go func() {
			errorsCh <- errors.New("positive worker count and non-nil handler required")
			close(errorsCh)
		}()
		return errorsCh
	}
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for job := range jobs {
				if err := runJob(job, handle); err != nil {
					errorsCh <- err
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(errorsCh)
	}()
	return errorsCh
}

// 7. Execute one task on each tick, with no initial burst.
func ProcessRateLimited(tasks []int, rate int, process func(int)) error {
	if rate <= 0 || rate > int(time.Second) {
		return errors.New("rate must be between 1 and 1000000000")
	}
	if process == nil {
		return errors.New("process function must not be nil")
	}
	// Round up so the interval never permits more than the requested rate.
	interval := (time.Second + time.Duration(rate) - 1) / time.Duration(rate)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for _, task := range tasks {
		<-ticker.C
		process(task)
	}
	return nil
}

// 8. Each goroutine owns one partial sum; Wait joins them before reading.
func ParallelSum(nums []int, workers int) int {
	if len(nums) == 0 {
		return 0
	}
	if workers <= 0 {
		workers = 1
	}
	if workers > len(nums) {
		workers = len(nums)
	}
	partial := make([]int, workers)
	var wg sync.WaitGroup
	wg.Add(workers)
	chunkSize, extra := len(nums)/workers, len(nums)%workers
	start := 0
	for i := 0; i < workers; i++ {
		end := start + chunkSize
		if i < extra {
			end++
		}
		go func(index, start, end int) {
			defer wg.Done()
			for _, value := range nums[start:end] {
				partial[index] += value
			}
		}(i, start, end)
		start = end
	}
	wg.Wait()
	total := 0
	for _, sum := range partial {
		total += sum
	}
	return total
}

func runQuestion2Examples() {
	fmt.Println("1. Sliding window:", SlidingWindowMax([]int{1, 3, -1, -3, 5, 3, 6, 7}, 4))
	fmt.Println("2. Anagrams:", GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	students := []Student{
		{Name: "Ali", ID: 1, Grades: map[string]float64{"Go": 80, "Math": 90}},
		{Name: "Mona", ID: 2, Grades: map[string]float64{"Go": 95, "Math": 99}},
	}
	top := TopStudent(students)
	fmt.Printf("3. Top student: %s, average: %.2f\n", top.Name, top.Average())
	people := ByAge{{Name: "Ali", Age: 25}, {Name: "Mona", Age: 20}, {Name: "Omar", Age: 30}}
	sort.Sort(people)
	fmt.Println("4. Sorted by age:", people)
	ranked := ByAge{{Name: "Omar", Age: 30}, {Name: "Mona", Age: 20}, {Name: "Ali", Age: 25}}
	SortRanked(ranked)
	fmt.Println("   Sorted through Ranker:", ranked)
	attempts := 0
	err := Retry(3, func() error {
		attempts++
		if attempts < 3 {
			return errors.New("temporary failure")
		}
		return nil
	})
	fmt.Printf("5. Retry: attempts=%d, error=%v\n", attempts, err)
	jobs := make(chan int, 5)
	for job := 1; job <= 5; job++ {
		jobs <- job
	}
	close(jobs)
	fmt.Println("6. Worker pool (error order can vary):")
	for err := range WorkerPool(jobs, 2, func(job int) error {
		if job == 2 {
			panic("unexpected failure")
		}
		if job == 4 {
			return errors.New("could not process task")
		}
		return nil
	}) {
		fmt.Println("  ", err)
	}
	tasks := make([]int, 20)
	for i := range tasks {
		tasks[i] = i + 1
	}
	fmt.Println("7. Processing 20 tasks at 10 per second:")
	if err := ProcessRateLimited(tasks, 10, func(task int) { fmt.Printf("%d ", task) }); err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	fmt.Println("8. Parallel sum:", ParallelSum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3))
}
