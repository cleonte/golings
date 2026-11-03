// concurrent17
// Make the tests pass!

package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// Cache is read by many goroutines and written to now and then.
type Cache struct {
	mu   sync.RWMutex
	data map[string]string

	// How many readers are inside Get, and the most there ever were. This only
	// serves the test: leave it as it is.
	active, most atomic.Int32
}

func (c *Cache) Get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	c.countReader()
	defer c.active.Add(-1)

	time.Sleep(50 * time.Millisecond) // a slow read, long enough for readers to overlap
	return c.data[key]
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
}

// countReader notes that one more reader is inside Get, and the highest number
// of readers seen at the same time.
func (c *Cache) countReader() {
	now := c.active.Add(1)
	for {
		seen := c.most.Load()
		if now <= seen || c.most.CompareAndSwap(seen, now) {
			return
		}
	}
}

func TestReadersCanRunTogether(t *testing.T) {
	cache := &Cache{data: map[string]string{"a": "1"}}
	start := make(chan struct{})

	var wg sync.WaitGroup
	for range 5 {
		wg.Go(func() {
			<-start // all the readers begin together
			if got := cache.Get("a"); got != "1" {
				t.Errorf("Get(a) = %q, want 1", got)
			}
		})
	}
	close(start)
	wg.Wait()

	if most := cache.most.Load(); most < 2 {
		t.Errorf("only one reader at a time (at most %d); reads should be able to run together", most)
	}
}

func TestAWriteIsSeenByTheNextRead(t *testing.T) {
	cache := &Cache{data: map[string]string{}}

	cache.Set("a", "2")

	if got := cache.Get("a"); got != "2" {
		t.Errorf("Get(a) = %q, want 2", got)
	}
}
