// concurrent15
// Make the tests pass!

package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"
)

// runAll runs every task in its own goroutine, with at most limit of them
// running at the same time, and returns when all of them are done.
func runAll(tasks []func(), limit int) {
	var wg sync.WaitGroup
	slots := make(chan struct{}, limit)

	for _, task := range tasks {
		slots <- struct{}{} // waits here while limit tasks are running
		wg.Go(func() {
			defer func() { <-slots }()
			task()
		})
	}

	wg.Wait()
}

func TestNoMoreThanTheLimitRunAtOnce(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var running, most atomic.Int32
		task := func() {
			now := running.Add(1)
			for {
				seen := most.Load()
				if now <= seen || most.CompareAndSwap(seen, now) {
					break
				}
			}
			time.Sleep(time.Second) // the work
			running.Add(-1)
		}
		tasks := make([]func(), 10)
		for i := range tasks {
			tasks[i] = task
		}
		start := time.Now()

		runAll(tasks, 3)

		if got := most.Load(); got != 3 {
			t.Errorf("up to %d tasks ran at the same time, want 3", got)
		}
		// 10 tasks, 3 at a time, a second each: 4 rounds.
		if elapsed := time.Since(start); elapsed != 4*time.Second {
			t.Errorf("took %v, want 4s", elapsed)
		}
	})
}
