// pitfalls5
// Make the tests pass!

package main_test

import (
	"fmt"
	"maps"
	"slices"
	"strings"
	"testing"
)

// describe lists the prices as "apple=3, banana=1" in alphabetical order.
func describe(prices map[string]int) string {
	var parts []string
	for _, name := range slices.Sorted(maps.Keys(prices)) {
		parts = append(parts, fmt.Sprintf("%s=%d", name, prices[name]))
	}
	return strings.Join(parts, ", ")
}

func TestDescribeIsAlwaysInOrder(t *testing.T) {
	prices := map[string]int{"apple": 3, "banana": 1, "cherry": 2, "date": 7, "elder": 4, "fig": 9, "grape": 5, "honey": 8, "iris": 6, "jam": 10}
	want := "apple=3, banana=1, cherry=2, date=7, elder=4, fig=9, grape=5, honey=8, iris=6, jam=10"

	// The order of a map is different from one range to the next: try many times.
	for range 50 {
		if got := describe(prices); got != want {
			t.Fatalf("describe gave\n%s\nwant\n%s", got, want)
		}
	}
}
