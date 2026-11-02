// maps_package3
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"testing"
)

// dropExpired removes the sessions that ended before now. The map holds the
// time each session ends.
func dropExpired(sessions map[string]int, now int) {
	// TODO: Delete the entries whose end time is before now, with one call to the maps package.
}

func TestDropExpired(t *testing.T) {
	sessions := map[string]int{"ana": 100, "carl": 250, "maria": 90, "john": 300}

	dropExpired(sessions, 200)

	if len(sessions) != 2 {
		t.Fatalf("%d sessions left, want 2: %v", len(sessions), sessions)
	}
	for _, name := range []string{"carl", "john"} {
		if _, ok := sessions[name]; !ok {
			t.Errorf("%s is still active, it should stay", name)
		}
	}
}
