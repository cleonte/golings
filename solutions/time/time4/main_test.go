// time4
// Make the tests pass!

package main_test

import (
	"testing"
	"time"
)

// startOfDay returns midnight at the start of the day t falls on, in t's own location.
func startOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
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
