// pitfalls6
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"testing"
)

// Notifier tells the goroutines that wait on Done that it is time to stop.
type Notifier struct {
	done chan struct{}
}

func NewNotifier() *Notifier {
	return &Notifier{done: make(chan struct{})}
}

func (n *Notifier) Done() <-chan struct{} {
	return n.done
}

// Stop closes Done. Several parts of the program can call it, so it has to be
// safe to call more than once.
func (n *Notifier) Stop() {
	// TODO: A second call panics, because the channel is already closed. Make sure it is closed only once.
	close(n.done)
}

func TestStopTwice(t *testing.T) {
	n := NewNotifier()

	n.Stop()
	n.Stop()

	select {
	case <-n.Done():
	default:
		t.Error("Done is not closed after Stop")
	}
}
