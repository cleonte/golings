// unsafe1
// Make the tests pass!

package main_test

import (
	"testing"
	"unsafe"
)

// Record is stored millions of times, so its size matters.
type Record struct {
	ID      int64
	Active  bool
	Deleted bool
}

func TestRecordHasNoWastedPadding(t *testing.T) {
	size := unsafe.Sizeof(Record{})

	if size > 16 {
		t.Errorf("a Record takes %d bytes, want 16 or fewer: the fields only need 10", size)
	}
}
