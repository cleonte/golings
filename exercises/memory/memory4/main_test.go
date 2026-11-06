// memory4
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"runtime"
	"testing"
	"time"
)

type resource struct{ data [1 << 20]byte }

// track arranges for closed to be called once r is no longer used and the
// garbage collector has found out.
func track(r *resource, closed chan<- struct{}) {
	// TODO: Nothing is arranged yet. runtime.AddCleanup(ptr, f, arg) runs f(arg) some time after ptr became unreachable. Close the channel in the cleanup (the argument can be the channel).
}

func create(closed chan<- struct{}) {
	track(&resource{}, closed) // nothing else refers to the resource after this
}

func TestTheCleanupRunsAfterTheResourceIsCollected(t *testing.T) {
	closed := make(chan struct{})
	create(closed)

	// Collection and cleanups are not instant: collect until the cleanup has run.
	timeout := time.After(3 * time.Second)
	for {
		runtime.GC()
		select {
		case <-closed:
			return
		case <-time.After(10 * time.Millisecond):
		case <-timeout:
			t.Fatal("the cleanup did not run: nothing is registered with runtime.AddCleanup")
		}
	}
}
