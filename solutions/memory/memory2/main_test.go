// memory2
// Make the tests pass!

package main_test

import (
	"bytes"
	"sync"
	"sync/atomic"
	"testing"
)

// greet builds "Hello, <name>!" in a buffer. Making a buffer for every call is
// wasteful when many calls run: take one from the pool and give it back.
func greet(pool *sync.Pool, name string) string {
	buf := pool.Get().(*bytes.Buffer)
	buf.Reset()
	defer pool.Put(buf)

	buf.WriteString("Hello, ")
	buf.WriteString(name)
	buf.WriteString("!")
	return buf.String()
}

// newPool makes a pool of buffers, and counts the buffers it had to make.
func newPool() (pool *sync.Pool, created *atomic.Int32) {
	created = new(atomic.Int32)
	pool = &sync.Pool{New: func() any {
		created.Add(1)
		return new(bytes.Buffer)
	}}
	return pool, created
}

func TestGreetStillGreets(t *testing.T) {
	pool, _ := newPool()

	for _, name := range []string{"Ada", "Alan", "Grace"} {
		if got, want := greet(pool, name), "Hello, "+name+"!"; got != want {
			t.Errorf("greet(%q) = %q, want %q", name, got, want)
		}
	}
}

func TestTheBuffersAreReused(t *testing.T) {
	pool, created := newPool()

	for range 100 {
		greet(pool, "Ada")
	}

	// The pool may drop a buffer now and then (at a garbage collection), so
	// the test only asks for far fewer than one buffer per call.
	if n := created.Load(); n == 0 || n > 50 {
		t.Errorf("the pool made %d buffers for 100 calls, want at least 1 and at most 50: take the buffers from the pool and put them back", n)
	}
}
