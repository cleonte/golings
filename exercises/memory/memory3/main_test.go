// memory3
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"runtime"
	"testing"
	"weak"
)

type blob struct{ data [1 << 20]byte }

// The weak package is imported for you; this line only keeps the import used until you use it yourself.
var _ = weak.Make[blob]

// cache remembers a blob without keeping it alive: when nobody else uses it,
// the garbage collector may free it, and the cache then has nothing.
type cache struct {
	// TODO: A *blob in a field is a strong reference: the blob lives as long as the cache does. Hold it with a weak pointer instead (the weak package), and make set and get use it.
	entry *blob
}

func (c *cache) set(b *blob) { c.entry = b }

// get returns the blob, or nil when it has been collected.
func (c *cache) get() *blob { return c.entry }

func fill(c *cache) {
	c.set(&blob{}) // the only strong reference is gone when fill returns
}

func TestTheBlobIsThereWhileSomeoneUsesIt(t *testing.T) {
	c := &cache{}
	b := &blob{}
	c.set(b)

	runtime.GC()

	if c.get() != b {
		t.Error("the blob was collected while it was still in use")
	}
	runtime.KeepAlive(b)
}

func TestTheCacheDoesNotKeepTheBlobAlive(t *testing.T) {
	c := &cache{}
	fill(c)

	runtime.GC()

	if c.get() != nil {
		t.Error("the blob is still there after a garbage collection: the cache holds it with a strong reference")
	}
}
