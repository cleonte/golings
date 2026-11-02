// slices_package1
// Make the tests pass!

package main_test

import (
	"cmp"
	"slices"
	"testing"
)

type Player struct {
	Name  string
	Score int
}

// sortNames puts the names in alphabetical order.
func sortNames(names []string) {
	slices.Sort(names)
}

// ranking puts the highest score first. Players with the same score go in
// alphabetical order.
func ranking(players []Player) {
	slices.SortFunc(players, func(a, b Player) int {
		return cmp.Or(
			cmp.Compare(b.Score, a.Score),
			cmp.Compare(a.Name, b.Name),
		)
	})
}

func TestSortNames(t *testing.T) {
	names := []string{"maria", "carl", "john", "ana"}

	sortNames(names)

	if want := []string{"ana", "carl", "john", "maria"}; !slices.Equal(names, want) {
		t.Errorf("sortNames gave %v, want %v", names, want)
	}
}

func TestRanking(t *testing.T) {
	players := []Player{{"carl", 7}, {"ana", 9}, {"maria", 7}, {"john", 3}, {"bea", 9}}

	ranking(players)

	want := []Player{{"ana", 9}, {"bea", 9}, {"carl", 7}, {"maria", 7}, {"john", 3}}
	if !slices.Equal(players, want) {
		t.Errorf("ranking gave %v, want %v", players, want)
	}
}
