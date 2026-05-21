package main

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"
)

type Job struct {
	ID int
}

type Result struct {
	JobID int
	Value string
	Err   error
}

// simulateWork returns an error for odd job IDs, to exercise the error path.
func simulateWork(ctx context.Context, job Job) (string, error) {
	select {
	case <-time.After(300 * time.Millisecond):
		// work completed
	case <-ctx.Done():
		// context cancelled mid-work
		return "", ctx.Err()
	}

	if job.ID%5 == 0 {
		return "", errors.New("job failed: divisible by 5")
	}
	return fmt.Sprintf("processed job %d", job.ID), nil
}

func worker(
	ctx context.Context,
	id int,
	jobs <-chan Job,
	results chan<- Result,
) {
	for {
		// Double-select: check cancellation first, before blocking on jobs.
		// Without this, Go randomly picks between ctx.Done() and jobs when
		// both are ready — cancellation might not win immediately.
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d: cancelled before picking up job\n", id)
			return
		default:
		}

		select {
		case <-ctx.Done():
			fmt.Printf("worker %d: cancelled while waiting for job\n", id)
			return

		case job, ok := <-jobs:
			// check if the jobs channel is closed
			if !ok {
				fmt.Printf("worker %d: jobs channel closed, exiting\n", id)
				return
			}

			value, err := simulateWork(ctx, job)

			select {
			case <-ctx.Done():
				fmt.Printf("worker %d: cancelled after picking up job %d, dropping result\n", id, job.ID)
				return
			case results <- Result{JobID: job.ID, Value: value, Err: err}:
			}
		}
	}
}

func main() {
	// WithTimeout to show a more realistic pattern than WithCancel alone.
	// 3 seconds should let most jobs through; cut it to 1s to see cancellation kick in.
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	const numJobs = 20
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	workerCount := runtime.NumCPU()
	fmt.Printf("starting %d workers\n", workerCount)

	// start workers
	var wg sync.WaitGroup
	for w := range workerCount {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, w+1, jobs, results)
		}()
	}

	// send jobs into the workers, and close jobs chan
	for i := range numJobs {
		jobs <- Job{ID: i + 1}
	}
	close(jobs)

	// start waiter so main can collect results concurrently
	go func() {
		wg.Wait()
		// all workers have exited, no more senders; close results chan
		close(results)
	}()

	// collect result
	var succeeded, failed int
	for r := range results {
		if r.Err != nil {
			fmt.Printf("job %d failed: %v\n", r.JobID, r.Err)
			failed++
		} else {
			fmt.Printf("job %d: %s\n", r.JobID, r.Value)
			succeeded++
		}
	}

	fmt.Printf("\ndone — %d succeeded, %d failed\n", succeeded, failed)

	// Check why we stopped: timeout, explicit cancel, or clean finish.
	if err := ctx.Err(); err != nil {
		fmt.Printf("context ended with %v\n", err)
	}
}

// The double-select is the most important new pattern. A lone `select` with `ctx.Done()` and `jobs`
// gives Go the freedom to choose randomly when both channels are ready. The leading select `{ case
// <-ctx.Done(): return; default: }` is non-blocking — it only fires if cancellation is already
// signalled, otherwise falls through immediately to the main select. This ensures cancellation is
// checked first on every iteration.

// `simulateWork` uses `select` inside the work itself so a long-running operation can be
// interrupted mid-flight rather than only between jobs. This is the right pattern for anything
// slow — a DB call, an HTTP request, file I/O. In practice you'd pass `ctx` directly to `db.
// QueryContext(ctx, ...)` or `http.NewRequestWithContext(ctx, ...)` and the library handles this
// for you.

// `WithTimeout` is more realistic than `WithCancel` alone for a server. In your URL shortener,
// every HTTP handler will create a child context from the request's context — Chi propagates it
// via `r.Context()` — so your DB calls automatically get cancelled if the client disconnects.

// The third `select` around `results <- ...` guards against a worker trying to send after the
// context is cancelled. Without it, a worker finishing its last job right as timeout hits could
// block forever on a send that nobody will ever receive.
// Try cutting the timeout to `1*time.Second` and watch the cancellation messages fire mid-pool.
