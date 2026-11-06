// time4
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"testing"
	"time"
)

// startOfDay returns midnight at the start of the day t falls on, in t's own location.
func startOfDay(t time.Time) time.Time {
	// TODO: Truncate counts from the zero time in UTC, so this gives midnight in UTC, not where t lives. Build the midnight from the year, month and day of t with time.Date, in t's location.
	return t.Truncate(24 * time.Hour)
}

func TestStartOfDay(t *testing.T) {
	cet := time.FixedZone("CET", 60*60)
	when := time.Date(2026, time.November, 5, 14, 30, 0, 0, cet)

	got := startOfDay(when)

	want := time.Date(2026, time.November, 5, 0, 0, 0, 0, cet)
	if !got.Equal(want) {
		t.Errorf("startOfDay = %v, want %v", got, want)
	}
}
