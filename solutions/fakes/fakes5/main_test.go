// fakes5
// Make the tests pass!

package main_test

import (
	"testing"
	"time"
)

// now is the clock of the program. A test replaces it to control the time.
var now = time.Now

// greeting depends on the hour.
func greeting() string {
	if now().Hour() < 12 {
		return "good morning"
	}
	return "good afternoon"
}

// setClock makes now return the given hour for the rest of the test.
func setClock(t *testing.T, hour int) {
	t.Helper()
	old := now
	t.Cleanup(func() { now = old })
	now = func() time.Time {
		return time.Date(2000, time.January, 1, hour, 0, 0, 0, time.UTC)
	}
}

func TestMorning(t *testing.T) {
	setClock(t, 9)

	if got := greeting(); got != "good morning" {
		t.Errorf("greeting at 9 = %q, want %q", got, "good morning")
	}
}

func TestTheClockIsBackAfterAnotherTest(t *testing.T) {
	// No setClock here: the real clock must be in place again.
	if got := now().Year(); got < 2026 {
		t.Errorf("the clock still shows the year %d: the earlier test did not put it back", got)
	}
}
