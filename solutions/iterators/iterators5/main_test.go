// iterators5
// Make the tests pass!

package main_test

import (
	"iter"
	"maps"
	"slices"
	"testing"
)

// evens passes on the even numbers of seq and drops the others.
func evens(seq iter.Seq[int]) iter.Seq[int] {
	return func(yield func(int) bool) {
		for n := range seq {
			if n%2 == 0 && !yield(n) {
				return
			}
		}
	}
}

// evenIDs returns the even keys of the map, in order.
func evenIDs(users map[int]string) []int {
	return slices.Sorted(evens(maps.Keys(users)))
}

func TestEvens(t *testing.T) {
	got := slices.Collect(evens(slices.Values([]int{1, 2, 3, 4, 5, 6})))

	if want := []int{2, 4, 6}; !slices.Equal(got, want) {
		t.Errorf("evens gave %v, want %v", got, want)
	}
}

func TestEvenIDs(t *testing.T) {
	users := map[int]string{1: "ana", 2: "carl", 3: "maria", 8: "john", 6: "bea"}

	got := evenIDs(users)

	if want := []int{2, 6, 8}; !slices.Equal(got, want) {
		t.Errorf("evenIDs gave %v, want %v", got, want)
	}
}
