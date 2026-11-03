// pitfalls7
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"sync"
	"testing"
)

// Counter counts, and is safe to use from several goroutines.
type Counter struct {
	mu sync.Mutex
	n  int
}

// Inc counts one more.
// TODO: Every call gets a copy of the counter, with a copy of the lock, and counts on the copy. Count on the real one.
func (c Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

func TestCounterCounts(t *testing.T) {
	var c Counter

	var wg sync.WaitGroup
	for range 100 {
		wg.Go(c.Inc)
	}
	wg.Wait()

	if got := c.Value(); got != 100 {
		t.Errorf("the counter is %d, want 100", got)
	}
}
