// sorting3
// Make the tests pass!

package main_test

import (
	"cmp"
	"maps"
	"slices"
	"strings"
	"testing"
)

// sortedNames returns the names in ages in alphabetical order, ignoring upper and lower case.
func sortedNames(ages map[string]int) []string {
	return slices.SortedFunc(maps.Keys(ages), func(a, b string) int {
		return cmp.Compare(strings.ToLower(a), strings.ToLower(b))
	})
}

func TestSortedNames(t *testing.T) {
	ages := map[string]int{"cherry": 3, "Banana": 2, "apple": 1}

	got := sortedNames(ages)

	want := []string{"apple", "Banana", "cherry"}
	if !slices.Equal(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}
