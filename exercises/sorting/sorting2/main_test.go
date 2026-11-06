// sorting2
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"slices"
	"testing"
)

// insertSorted puts v into the sorted slice s and keeps it sorted.
func insertSorted(s []int, v int) []int {
	// TODO: v goes at the end, wherever it belongs. Find the position where v keeps the slice sorted, and insert it there.
	return append(s, v)
}

func TestInsertSorted(t *testing.T) {
	for _, tc := range []struct {
		v    int
		want []int
	}{
		{4, []int{1, 3, 4, 5, 7}},
		{0, []int{0, 1, 3, 5, 7}},
		{9, []int{1, 3, 5, 7, 9}},
	} {
		got := insertSorted([]int{1, 3, 5, 7}, tc.v)

		if !slices.Equal(got, tc.want) {
			t.Errorf("inserting %d gave %v, want %v", tc.v, got, tc.want)
		}
	}
}
