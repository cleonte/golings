// bench2
// Make the tests pass!

package main_test

import (
	"testing"
)

// squares returns the squares of 0, 1, ... n-1.
func squares(n int) []int {
	result := make([]int, 0, n)
	for i := range n {
		result = append(result, i*i)
	}
	return result
}

var sinkInts []int

func TestSquaresAllocatesOnce(t *testing.T) {
	got := squares(1000)
	if len(got) != 1000 || got[999] != 999*999 {
		t.Fatalf("squares(1000) has %d numbers and ends with %d, want 1000 numbers ending with %d", len(got), got[len(got)-1], 999*999)
	}

	allocs := testing.AllocsPerRun(10, func() { sinkInts = squares(1000) })

	if allocs > 1 {
		t.Errorf("squares allocates more than once per call, want exactly 1")
	}
}
