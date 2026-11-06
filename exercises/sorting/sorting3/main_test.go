// sorting3
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"maps"
	"slices"
	"testing"
)

// sortedNames returns the names in ages in alphabetical order, ignoring upper and lower case.
func sortedNames(ages map[string]int) []string {
	// TODO: Capital letters sort before small ones, so "Banana" comes before "apple". Collect the keys with a comparison that ignores the case instead.
	return slices.Sorted(maps.Keys(ages))
}

func TestSortedNames(t *testing.T) {
	ages := map[string]int{"cherry": 3, "Banana": 2, "apple": 1}

	got := sortedNames(ages)

	want := []string{"apple", "Banana", "cherry"}
	if !slices.Equal(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}
