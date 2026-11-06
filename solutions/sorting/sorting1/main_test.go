// sorting1
// Make the tests pass!

package main_test

import (
	"cmp"
	"slices"
	"testing"
)

type person struct{ Name, City string }

// sortPeople orders people by city, and by name inside a city.
func sortPeople(people []person) {
	slices.SortFunc(people, func(a, b person) int {
		return cmp.Or(
			cmp.Compare(a.City, b.City),
			cmp.Compare(a.Name, b.Name),
		)
	})
}

func TestSortPeople(t *testing.T) {
	people := []person{{"Zoe", "Oslo"}, {"Ann", "Rome"}, {"Bob", "Oslo"}, {"Cid", "Rome"}}

	sortPeople(people)

	want := []person{{"Bob", "Oslo"}, {"Zoe", "Oslo"}, {"Ann", "Rome"}, {"Cid", "Rome"}}
	if !slices.Equal(people, want) {
		t.Errorf("got %v, want %v", people, want)
	}
}
