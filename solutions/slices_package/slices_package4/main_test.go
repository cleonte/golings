// slices_package4
// Make the tests pass!

package main_test

import (
	"slices"
	"testing"
)

// insertAt returns the names with name added at position i.
func insertAt(names []string, i int, name string) []string {
	return slices.Insert(names, i, name)
}

// removeRange returns the names without the ones from position i up to, but
// not including, position j.
func removeRange(names []string, i, j int) []string {
	return slices.Delete(names, i, j)
}

// withoutRepeats returns a sorted slice with every value only once.
func withoutRepeats(sorted []int) []int {
	return slices.Compact(sorted)
}

func TestInsertAt(t *testing.T) {
	got := insertAt([]string{"ana", "maria"}, 1, "carl")

	if want := []string{"ana", "carl", "maria"}; !slices.Equal(got, want) {
		t.Errorf("insertAt gave %v, want %v", got, want)
	}
}

func TestRemoveRange(t *testing.T) {
	got := removeRange([]string{"ana", "carl", "john", "maria"}, 1, 3)

	if want := []string{"ana", "maria"}; !slices.Equal(got, want) {
		t.Errorf("removeRange gave %v, want %v", got, want)
	}
}

func TestWithoutRepeats(t *testing.T) {
	got := withoutRepeats([]int{1, 1, 2, 3, 3, 3, 4})

	if want := []int{1, 2, 3, 4}; !slices.Equal(got, want) {
		t.Errorf("withoutRepeats gave %v, want %v", got, want)
	}
}
