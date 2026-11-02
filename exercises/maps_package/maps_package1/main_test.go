// maps_package1
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"slices"
	"testing"
)

// names returns the names in the map in alphabetical order.
func names(ages map[string]int) []string {
	// TODO: Collect the keys of the map, sorted: the maps and slices packages have what you need.
	return nil
}

func TestNames(t *testing.T) {
	ages := map[string]int{"maria": 31, "ana": 28, "john": 45, "carl": 19}

	got := names(ages)

	if want := []string{"ana", "carl", "john", "maria"}; !slices.Equal(got, want) {
		t.Errorf("names gave %v, want %v", got, want)
	}
}
