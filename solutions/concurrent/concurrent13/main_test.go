//go:build go1.27

// concurrent13
// Make the tests pass!

package main_test

import (
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"
)

func TestEveryWorkerHasFinished(t *testing.T) {
	// Which goroutine runs first when several wake up together is random, so
	// the test repeats the check.
	for range 200 {
		synctest.Test(t, func(t *testing.T) {
			var finished atomic.Int32
			for range 5 {
				go func() {
					time.Sleep(time.Minute)
					finished.Add(1)
				}()
			}

			synctest.Sleep(time.Minute)

			if got := finished.Load(); got != 5 {
				t.Fatalf("%d of the 5 workers have finished, want all of them", got)
			}
		})
	}
}
