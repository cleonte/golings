// concurrent17
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"sync"
	"testing"
	"time"
)

// Cache is read by many goroutines and written to now and then.
type Cache struct {
	mu   sync.Mutex
	data map[string]string
}

func (c *Cache) Get(key string) string {
	// TODO: Reads don't need to wait for each other, only writes have to be alone. Use a lock that lets many readers in, and take it for reading here.
	c.mu.Lock()
	defer c.mu.Unlock()

	time.Sleep(100 * time.Millisecond) // a slow read
	return c.data[key]
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
}

func TestReadersDoNotWaitForEachOther(t *testing.T) {
	cache := &Cache{data: map[string]string{"a": "1"}}
	start := time.Now()

	var wg sync.WaitGroup
	for range 5 {
		wg.Go(func() {
			if got := cache.Get("a"); got != "1" {
				t.Errorf("Get(a) = %q, want 1", got)
			}
		})
	}
	wg.Wait()

	// Together the five reads take about 100ms, one after the other 500ms.
	if elapsed := time.Since(start); elapsed > 300*time.Millisecond {
		t.Errorf("5 reads took %v, want about 100ms: they should run together", elapsed.Round(10*time.Millisecond))
	}
}

func TestAWriteIsSeenByTheNextRead(t *testing.T) {
	cache := &Cache{data: map[string]string{}}

	cache.Set("a", "2")

	if got := cache.Get("a"); got != "2" {
		t.Errorf("Get(a) = %q, want 2", got)
	}
}
