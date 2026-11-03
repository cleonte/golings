// pitfalls2
// Make the tests pass!

package main_test

import "testing"

type Player struct {
	Name  string
	Score int
}

// addBonus gives every player 10 more points.
func addBonus(players []Player) {
	for i := range players {
		players[i].Score += 10
	}
}

func TestAddBonus(t *testing.T) {
	players := []Player{{"ana", 5}, {"carl", 12}}

	addBonus(players)

	if players[0].Score != 15 || players[1].Score != 22 {
		t.Errorf("scores are %d and %d, want 15 and 22", players[0].Score, players[1].Score)
	}
}
