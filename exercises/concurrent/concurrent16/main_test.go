// concurrent16
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"
)

// runTasks runs the tasks at the same time and waits for all of them. It
// returns the first error that a task returned, and the other tasks are told to
// stop by cancelling the context they were given.
func runTasks(tasks []func(ctx context.Context) error) error {
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	var wg sync.WaitGroup
	var firstErr error

	for _, task := range tasks {
		wg.Go(func() {
			// TODO: Run the task with ctx. When it fails, keep the first error only, and cancel the context so the other tasks stop.
			task(ctx)
		})
	}

	wg.Wait()
	return firstErr
}

func TestTheFirstErrorWinsAndTheOthersStop(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		var stopped atomic.Bool
		tasks := []func(ctx context.Context) error{
			func(ctx context.Context) error { // fails after a second
				time.Sleep(time.Second)
				return errors.New("boom")
			},
			func(ctx context.Context) error { // would take an hour
				select {
				case <-time.After(time.Hour):
					return nil
				case <-ctx.Done():
					stopped.Store(true)
					return ctx.Err()
				}
			},
			func(ctx context.Context) error { // would fail later
				select {
				case <-time.After(2 * time.Second):
					return errors.New("later")
				case <-ctx.Done():
					return ctx.Err()
				}
			},
		}
		start := time.Now()

		err := runTasks(tasks)

		if err == nil || err.Error() != "boom" {
			t.Errorf("err = %v, want the first error, boom", err)
		}
		if !stopped.Load() {
			t.Error("the slow task was not told to stop")
		}
		if elapsed := time.Since(start); elapsed != time.Second {
			t.Errorf("took %v, want 1s: the others should stop when one fails", elapsed)
		}
	})
}

func TestNoErrorWhenAllTasksWork(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		ok := func(ctx context.Context) error { return nil }

		if err := runTasks([]func(ctx context.Context) error{ok, ok}); err != nil {
			t.Errorf("err = %v, want nil", err)
		}
	})
}
