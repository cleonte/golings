// pitfalls5
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"fmt"
	"strings"
	"testing"
)

// describe lists the prices as "apple=3, banana=1" in alphabetical order.
func describe(prices map[string]int) string {
	var parts []string
	// TODO: Going through a map gives a different order every time. The list has to be in alphabetical order, always.
	for name, price := range prices {
		parts = append(parts, fmt.Sprintf("%s=%d", name, price))
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
