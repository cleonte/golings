// iterators4
// Make the tests pass!

package main_test

import (
	"iter"
	"slices"
	"testing"
)

// naturals gives 0, 1, 2, 3... and never ends by itself.
func naturals() iter.Seq[int] {
	return func(yield func(int) bool) {
		for n := 0; ; n++ {
			if !yield(n) {
				return
			}
		}
	}
}

func TestNaturalsStopWhenTheLoopBreaks(t *testing.T) {
	var got []int
	for n := range naturals() {
		if n == 3 {
			break
		}
		got = append(got, n)
	}

	if want := []int{0, 1, 2}; !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
