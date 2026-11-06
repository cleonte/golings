// time2
// Make the tests pass!

package main_test

import (
	"testing"
	"time"
)

// formatStamp writes a time like "Thu, 05 Nov 2026 14:30".
func formatStamp(t time.Time) string {
	return t.Format("Mon, 02 Jan 2006 15:04")
}

func TestFormatStamp(t *testing.T) {
	when := time.Date(2026, time.November, 5, 14, 30, 0, 0, time.UTC)

	if got, want := formatStamp(when), "Thu, 05 Nov 2026 14:30"; got != want {
		t.Errorf("formatStamp = %q, want %q", got, want)
	}
}
