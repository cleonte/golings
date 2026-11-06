// unsafe3
// Make the tests pass!

package main_test

import (
	"testing"
	"unsafe"
)

// asBytes returns the memory of nums as bytes, without copying it.
func asBytes(nums []uint32) []byte {
	if len(nums) == 0 {
		return nil
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(nums))), len(nums)*int(unsafe.Sizeof(nums[0])))
}

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
