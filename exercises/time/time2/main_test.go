// time2
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"testing"
	"time"
)

// formatStamp writes a time like "Thu, 05 Nov 2026 14:30".
func formatStamp(t time.Time) string {
	// TODO: Go has no YYYY-MM-DD. A layout is the reference time, Mon Jan 2 15:04:05 MST 2006, written the way you want the result to look. Write the layout for the example above.
	return t.Format("DDD, DD MMM YYYY HH:mm")
}

func TestFormatStamp(t *testing.T) {
	when := time.Date(2026, time.November, 5, 14, 30, 0, 0, time.UTC)

	if got, want := formatStamp(when), "Thu, 05 Nov 2026 14:30"; got != want {
		t.Errorf("formatStamp = %q, want %q", got, want)
	}
}
