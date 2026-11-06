// time3
// Make the tests pass!

package main_test

import (
	"testing"
	"time"
)

// meetingUTC reads a meeting time written in the local time of zone, and returns it in UTC.
func meetingUTC(text string, zone *time.Location) (time.Time, error) {
	local, err := time.ParseInLocation("2006-01-02 15:04", text, zone)
	if err != nil {
		return time.Time{}, err
	}
	return local.UTC(), nil
}

func TestMeetingUTC(t *testing.T) {
	cet := time.FixedZone("CET", 60*60)

	got, err := meetingUTC("2026-03-08 09:30", cet)
	if err != nil {
		t.Fatal(err)
	}

	want := time.Date(2026, time.March, 8, 8, 30, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Errorf("meetingUTC = %v, want %v", got, want)
	}
	if got.Location() != time.UTC {
		t.Errorf("the result is in %v, want UTC", got.Location())
	}
}

func TestMeetingUTCRejectsADateThatDoesNotExist(t *testing.T) {
	if _, err := meetingUTC("2026-02-30 10:00", time.UTC); err == nil {
		t.Error("meetingUTC accepted the 30th of February")
	}
}
