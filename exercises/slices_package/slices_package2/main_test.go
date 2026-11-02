// slices_package2
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"testing"
)

// find looks for target in the sorted slice. It returns its position and
// whether it is there; if it is not, the position where it would go.
func find(sorted []int, target int) (int, bool) {
	// TODO: The slice is sorted: use the slices package to search it.
	return 0, false
}

func TestFind(t *testing.T) {
	sorted := []int{1, 3, 5, 7, 9, 11}

	tests := []struct {
		target    int
		wantIndex int
		wantFound bool
	}{
		{target: 7, wantIndex: 3, wantFound: true},
		{target: 1, wantIndex: 0, wantFound: true},
		{target: 4, wantIndex: 2, wantFound: false},
		{target: 0, wantIndex: 0, wantFound: false},
		{target: 12, wantIndex: 6, wantFound: false},
	}

	for _, tc := range tests {
		index, found := find(sorted, tc.target)
		if index != tc.wantIndex || found != tc.wantFound {
			t.Errorf("find(%d) = %d, %t, want %d, %t", tc.target, index, found, tc.wantIndex, tc.wantFound)
		}
	}
}
