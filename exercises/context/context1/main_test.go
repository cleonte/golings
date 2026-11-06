// context1
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"context"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"
)

// work does up to steps units of work, a second each, and counts the ones it finished.
func work(ctx context.Context, steps int, done *atomic.Int32) {
	for range steps {
		// TODO: Sleep can't be interrupted, so cancelling ctx changes nothing. Wait for either the second to pass or ctx to be done, and stop in the second case.
		time.Sleep(time.Second)
		done.Add(1)
	}
}

func TestWorkStopsWhenCancelled(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		var done atomic.Int32
		finished := make(chan struct{})
		go func() {
			defer close(finished)
			work(ctx, 10, &done)
		}()

		time.Sleep(2*time.Second + time.Millisecond)
		cancel()
		synctest.Wait()

		select {
		case <-finished:
		default:
			t.Error("work is still running after its context was cancelled")
		}
		if got := done.Load(); got != 2 {
			t.Errorf("work finished %d steps before it was cancelled, want 2", got)
		}
		<-finished // do not leave the worker behind, whatever it did
	})
}
