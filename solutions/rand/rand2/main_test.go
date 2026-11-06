// rand2
// Make the tests pass!

package main_test

import (
	"math/rand/v2"
	"slices"
	"testing"
)

// shuffled returns the items in a random order that depends only on the seed:
// the same seed always gives the same order.
func shuffled(items []int, seed uint64) []int {
	r := rand.New(rand.NewPCG(seed, seed))
	result := slices.Clone(items)
	r.Shuffle(len(result), func(i, j int) { result[i], result[j] = result[j], result[i] })
	return result
}

func TestTheSameSeedGivesTheSameOrder(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	first := shuffled(items, 42)
	second := shuffled(items, 42)

	if !slices.Equal(first, second) {
		t.Errorf("the same seed gave %v and then %v", first, second)
	}
}

func TestADifferentSeedGivesAnotherOrder(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	if slices.Equal(shuffled(items, 1), shuffled(items, 2)) {
		t.Error("two different seeds gave the same order")
	}
}
