// iterators1
// Make the tests pass!

package main_test

import (
	"iter"
	"slices"
	"testing"
)

// countdown gives the numbers from `from` down to 1, one at a time.
func countdown(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for n := from; n > 0; n-- {
			if !yield(n) {
				return
			}
		}
	}
}

func TestCountdown(t *testing.T) {
	got := slices.Collect(countdown(3))

	if want := []int{3, 2, 1}; !slices.Equal(got, want) {
		t.Errorf("countdown(3) gave %v, want %v", got, want)
	}
}

func TestCountdownInALoop(t *testing.T) {
	var got []int
	for n := range countdown(2) {
		got = append(got, n)
	}

	if want := []int{2, 1}; !slices.Equal(got, want) {
		t.Errorf("the loop saw %v, want %v", got, want)
	}
}
