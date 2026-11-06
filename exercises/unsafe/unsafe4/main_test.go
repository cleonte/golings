// unsafe4
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"testing"
	"unsafe"
)

// at returns nums[i] by moving a pointer, without the bounds check.
func at(nums []int32, i int) int32 {
	start := unsafe.Pointer(unsafe.SliceData(nums))
	// TODO: unsafe.Add moves by bytes, not by elements, so adding i lands in the middle of an element. Move by i times the size of one element.
	return *(*int32)(unsafe.Add(start, i))
}

func TestAtFindsTheElement(t *testing.T) {
	nums := []int32{10, 20, 30, 40}

	for i, want := range nums {
		if got := at(nums, i); got != want {
			t.Errorf("at(nums, %d) = %d, want %d", i, got, want)
		}
	}
}
