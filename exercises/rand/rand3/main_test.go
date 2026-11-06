// rand3
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"math/rand/v2"
	"testing"
	"time"
)

// jitter returns a random wait of up to 5 seconds, not including 5 seconds: any
// duration, nanoseconds included.
func jitter() time.Duration {
	// TODO: This only gives whole seconds. The generic function rand.N takes a number of any integer type, a Duration too, and returns one below it.
	return time.Duration(rand.IntN(5)) * time.Second
}

func TestJitter(t *testing.T) {
	var fractions int
	for range 200 {
		d := jitter()
		if d < 0 || d >= 5*time.Second {
			t.Fatalf("jitter() = %v, want at least 0 and less than 5s", d)
		}
		if d%time.Second != 0 {
			fractions++
		}
	}

	if fractions == 0 {
		t.Error("jitter() only returned whole seconds in 200 tries: it should be any duration below 5s")
	}
}
