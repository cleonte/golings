// unsafe1
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"testing"
	"unsafe"
)

// Record is stored millions of times, so its size matters.
type Record struct {
	// TODO: The compiler pads every field up to its alignment: an int64 starts at a multiple of 8, so the bools around it cost 7 bytes each. Put the fields in an order that wastes less, largest first.
	Active  bool
	ID      int64
	Deleted bool
}

func TestRecordHasNoWastedPadding(t *testing.T) {
	size := unsafe.Sizeof(Record{})

	if size > 16 {
		t.Errorf("a Record takes %d bytes, want 16 or fewer: the fields only need 10", size)
	}
}
