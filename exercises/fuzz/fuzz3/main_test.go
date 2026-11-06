// fuzz3
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"testing"
)

// Clamp limits n to the range from lo to hi.
func Clamp(n, lo, hi int) int {
	return max(lo, min(n, hi))
}

func FuzzClamp(f *testing.F) {
	// TODO: The seeds are text, and the fuzz function takes three ints, which the vet check of go test reports. f.Add takes the same types, in the same order, as the parameters after t.
	f.Add("5", "0", "10")
	f.Add("-3", "0", "10")
	f.Add("99", "0", "10")
	f.Fuzz(func(t *testing.T, n, lo, hi int) {
		if lo > hi {
			t.Skip()
		}
		if got := Clamp(n, lo, hi); got < lo || got > hi {
			t.Errorf("Clamp(%d, %d, %d) = %d, outside the range", n, lo, hi, got)
		}
	})
}
