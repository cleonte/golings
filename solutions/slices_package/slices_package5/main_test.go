// slices_package5
// Make the tests pass!

package main_test

import (
	"slices"
	"testing"
)

// sorted returns the scores in order, and leaves the caller's slice as it was.
func sorted(scores []int) []int {
	copied := slices.Clone(scores)
	slices.Sort(copied)
	return copied
}

func TestSortedLeavesTheOriginalAlone(t *testing.T) {
	scores := []int{30, 10, 20}

	got := sorted(scores)

	if want := []int{10, 20, 30}; !slices.Equal(got, want) {
		t.Errorf("sorted gave %v, want %v", got, want)
	}
	if want := []int{30, 10, 20}; !slices.Equal(scores, want) {
		t.Errorf("the original slice was changed to %v, want it to stay %v", scores, want)
	}
}
