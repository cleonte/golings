// maps_package1
// Make the tests pass!

package main_test

import (
	"maps"
	"slices"
	"testing"
)

// names returns the names in the map in alphabetical order.
func names(ages map[string]int) []string {
	return slices.Sorted(maps.Keys(ages))
}

func TestNames(t *testing.T) {
	ages := map[string]int{"maria": 31, "ana": 28, "john": 45, "carl": 19}

	got := names(ages)

	if want := []string{"ana", "carl", "john", "maria"}; !slices.Equal(got, want) {
		t.Errorf("names gave %v, want %v", got, want)
	}
}
