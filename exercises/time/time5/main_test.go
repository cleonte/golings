// time5
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"testing"
	"time"
)

// sameInstant reports whether a and b are the same moment, wherever each is written.
func sameInstant(a, b time.Time) bool {
	// TODO: == compares the location (and more) as well as the moment. Use the method that asks whether two times are the same instant.
	return a == b
}

func TestSameInstant(t *testing.T) {
	noon := time.Date(2026, time.November, 5, 12, 0, 0, 0, time.UTC)
	sameNoonInCET := noon.In(time.FixedZone("CET", 60*60)) // 13:00 there

	if !sameInstant(noon, sameNoonInCET) {
		t.Errorf("%v and %v are the same instant, sameInstant said they are not", noon, sameNoonInCET)
	}
}

func TestDifferentInstants(t *testing.T) {
	noon := time.Date(2026, time.November, 5, 12, 0, 0, 0, time.UTC)
	later := noon.Add(time.Nanosecond)

	if sameInstant(noon, later) {
		t.Error("two different instants are reported as the same")
	}
}
