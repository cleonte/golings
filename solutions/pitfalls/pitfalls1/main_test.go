// pitfalls1
// Make the tests pass!

package main_test

import (
	"slices"
	"testing"
)

// withExtra returns the numbers of base followed by extra, as a new slice. The
// base must stay as it was, and so must the results of other calls with it.
func withExtra(base []int, extra int) []int {
	result := slices.Clone(base)
	return append(result, extra)
}

func TestWithExtraDoesNotShareTheArray(t *testing.T) {
	base := make([]int, 3, 10) // room for more
	copy(base, []int{1, 2, 3})

	a := withExtra(base, 4)
	b := withExtra(base, 5)

	if want := []int{1, 2, 3, 4}; !slices.Equal(a, want) {
		t.Errorf("a = %v, want %v", a, want)
	}
	if want := []int{1, 2, 3, 5}; !slices.Equal(b, want) {
		t.Errorf("b = %v, want %v", b, want)
	}
}
