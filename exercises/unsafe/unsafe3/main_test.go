// unsafe3
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"testing"
	"unsafe"
)

// asBytes returns the memory of nums as bytes, without copying it.
func asBytes(nums []uint32) []byte {
	// TODO: This returns new, empty bytes. Build a slice over the memory of nums: unsafe.Slice takes a *byte (convert unsafe.SliceData(nums) with unsafe.Pointer) and a length in bytes, which is len(nums) times unsafe.Sizeof(nums[0]).
	return make([]byte, len(nums)*4)
}

var _ unsafe.Pointer

func TestAsBytesSharesTheMemory(t *testing.T) {
	nums := []uint32{1, 2}
	b := asBytes(nums)

	if len(b) != 8 {
		t.Fatalf("len(asBytes(nums)) = %d, want 8", len(b))
	}
	for i := range b {
		b[i] = 0xFF
	}
	for i, n := range nums {
		if n != 0xFFFFFFFF {
			t.Errorf("nums[%d] = %#x after filling the bytes with 0xFF, want 0xffffffff: the bytes are a copy", i, n)
		}
	}
}
