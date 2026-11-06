// context6
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"context"
	"sync/atomic"
	"testing"
	"testing/synctest"
)

type conn struct{ closed atomic.Bool }

func (c *conn) Close() { c.closed.Store(true) }

// closeWhenDone closes c as soon as ctx is done, without a goroutine of its own
// waiting for that. The function it returns takes the closing back: it reports
// true if it did that before c was closed.
func closeWhenDone(ctx context.Context, c *conn) (stop func() bool) {
	// TODO: Nothing closes c yet. context.AfterFunc runs a function once ctx is done, and returns the stop function this one has to return.
	return func() bool { return false }
}

func TestTheConnectionIsClosedWhenTheContextEnds(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		c := &conn{}
		closeWhenDone(ctx, c)

		cancel()
		synctest.Wait()

		if !c.closed.Load() {
			t.Error("the connection is still open after the context was cancelled")
		}
	})
}

func TestStoppingKeepsTheConnectionOpen(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		c := &conn{}
		stop := closeWhenDone(ctx, c)

		if !stop() {
			t.Error("stop() = false, want true: it was called before the context ended")
		}
		cancel()
		synctest.Wait()

		if c.closed.Load() {
			t.Error("the connection was closed although stop had been called")
		}
	})
}
