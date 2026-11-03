// concurrent18
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"
)

// Election lets many goroutines try to become the leader. One of them, and
// only one, wins.
type Election struct {
	taken atomic.Bool
}

// TryLead reports whether the caller became the leader.
func (e *Election) TryLead() bool {
	// TODO: Checking the flag and setting it are two steps, and other goroutines get in between. Make them one step.
	if e.taken.Load() {
		return false
	}
	time.Sleep(time.Millisecond) // some work before taking the role
	e.taken.Store(true)
	return true
}

func TestExactlyOneGoroutineWins(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var election Election
		var winners atomic.Int32

		var wg sync.WaitGroup
		for range 10 {
			wg.Go(func() {
				if election.TryLead() {
					winners.Add(1)
				}
			})
		}
		wg.Wait()

		if got := winners.Load(); got != 1 {
			t.Errorf("%d goroutines became the leader, want exactly 1", got)
		}
	})
}
