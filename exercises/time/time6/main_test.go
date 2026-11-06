// time6
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"
)

// debouncer runs fn once things have been quiet for wait: every Trigger
// postpones the call, so a burst of triggers ends in a single call.
type debouncer struct {
	wait  time.Duration
	fn    func()
	timer *time.Timer
}

func (d *debouncer) Trigger() {
	// TODO: Every call starts a new timer, so each trigger of a burst runs fn. Keep one timer: start it on the first trigger, and with later ones move it back to a full wait.
	time.AfterFunc(d.wait, d.fn)
}

func TestABurstIsOneCall(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var calls atomic.Int32
		d := &debouncer{wait: 100 * time.Millisecond, fn: func() { calls.Add(1) }}

		for range 3 {
			d.Trigger()
			time.Sleep(50 * time.Millisecond)
		}
		if got := calls.Load(); got != 0 {
			t.Errorf("fn ran %d times during the burst, want 0: every trigger should postpone it", got)
		}

		time.Sleep(time.Second)
		synctest.Wait()
		if got := calls.Load(); got != 1 {
			t.Errorf("fn ran %d times in all, want 1", got)
		}
	})
}
