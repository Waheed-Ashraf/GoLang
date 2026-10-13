# Go exam — Questions 2 and 3

Question 3's mini project is in `question3.go`; all eight Question 2 solutions
are in `question2.go`. The Question 1 notes in `main.go` are preserved.

Run from this folder:

```powershell
go run .
go run . -workers=5
go run . -question=2
go test ./...
```

If the default Go build cache is inaccessible, set a local cache first:

```powershell
$env:GOCACHE = Join-Path (Get-Location) '.go-cache'
```

## Question 3: task worker pool

`go run .` runs Question 3. `Task.Name` is the type key (`email`, `resize`, or
`report`) in a `map[string]Handler`. Each handler validates its payload and
simulates work. Payload flags `temporary_failure`, `permanent_failure`, and
`panic` demonstrate the different failure paths without external services.

Workers read a task channel and call the matching handler. `RetryableError`
failures get up to **three retries after the initial attempt** (four attempts
total), waiting 20ms, 40ms, and 60ms between attempts. `errors.As` detects
retryable errors through `%w` wrappers. Other errors and recovered `PanicError`
failures are not retried. Each attempt has its own deferred recovery, so a panic
does not stop the worker from taking another task.

A mutex protects both the results slice and statistics. The demo prints each
result and a summary: **4 succeeded, 4 failed, 1 retried task, 1 panicked task,
and 3 retry attempts**. Panicked tasks are included in failed tasks. Average
duration is per recorded task, including retries and backoff; it is not the
whole pool's wall-clock duration divided by the task count.

The demo submits tasks in descending priority order. Concurrent workers may
finish in any order; the printed results are sorted by ID. The pool consumes
the channel in its existing order rather than implementing a priority queue.

Closing the input channel lets all queued work finish. **Ctrl+C** cancels the
context: workers stop accepting tasks, active handlers observe cancellation,
retry waits stop, and the pool waits for all workers before printing its partial
summary. Tasks not started at cancellation are not included in the results.
Custom handlers must honor the context for prompt shutdown. Keep payload maps
and the handler registry immutable while workers run, and protect any shared
state inside custom handlers.

`question3_test.go` verifies successful retries, the retry limit, error wrapping,
panic recovery followed by more work, concurrent statistics, cancellation,
handler validation, and invalid worker counts. If a C compiler is available,
run `go test -race ./...` to check concurrent access too.

## Question 2: how the solutions work

1. **Sliding window maximum:** Store indices in decreasing value order. Remove indices outside the window and smaller values at the back. The front is the current maximum. The example produces `[3 3 5 5 6 7]`. Time: O(n). Invalid window sizes return an empty slice.
2. **Group anagrams:** Sort each word's letters to make a map key. Words with the same key enter the same group. Output: `[[eat tea ate] [tan nat] [bat]]`. Groups and their words retain first appearance order; letter comparisons are case sensitive.
3. **Students:** `Average` divides the sum of grades by their count, returning zero when there are no grades. `TopStudent` returns a pointer into the original slice, or nil for an empty slice. The first student wins a tie.
4. **Interfaces:** `ByAge` implements `Len`, `Less`, and `Swap` for `sort.Interface`. The custom `Ranker` exposes `Len`, `Rank`, and `Swap`. An adapter provides `Less` so any collection implementing `Ranker` can use `SortRanked`. Lower ranks come first; the example ranks people by age.
5. **Retry:** Try at most `times` times and stop immediately on success. Wait 100ms, 200ms, 400ms, and so on between failures, capped at two seconds. After the final failure, wrap its error with `fmt.Errorf("after %d attempts: %w", times, lastErr)`. `%w` preserves the original error for `errors.Is` and `errors.As`. Nonpositive attempt counts return an error without calling the function.
6. **Panic recovery:** An unrecovered panic in a goroutine crashes the entire program. Each worker invokes a job through a helper with deferred recovery, in the worker's own goroutine. The helper converts panics to errors so the worker can continue. Close the jobs channel when finished sending, and drain the error channel until it closes. Error ordering can vary between workers.
7. **Rate limiter:** Wait for a `time.Ticker` tick before processing each task. The example processes all 20 tasks at 10 operations per second, taking about two seconds. Stop the ticker with `defer` when finished.
8. **Parallel sum:** Divide the slice into balanced contiguous chunks. Each goroutine writes to its own partial sum; `sync.WaitGroup` waits for all workers before combining them. The example returns `55`. Empty input returns zero, nonpositive worker counts use one worker, and worker counts above the slice length are reduced to that length.

