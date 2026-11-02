// builtins2
// Make the tests pass!

// I AM NOT DONE
package main_test

import "testing"

// forgetAll empties the cache. Other parts of the program hold the same map, so
// it has to be emptied, not replaced.
func forgetAll(cache map[string]int) {
	// TODO: Remove every entry from the map. Go 1.21 has a built-in for this.
}

func TestForgetAll(t *testing.T) {
	cache := map[string]int{"a": 1, "b": 2, "c": 3}
	sameCache := cache

	forgetAll(cache)

	if len(sameCache) != 0 {
		t.Errorf("the cache still has %d entries, want 0", len(sameCache))
	}

	cache["d"] = 4
	if sameCache["d"] != 4 {
		t.Error("the cache should still be a usable map, shared with its other holders")
	}
}
