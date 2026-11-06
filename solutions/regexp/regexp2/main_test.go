// regexp2
// Make the tests pass!

package main_test

import (
	"regexp"
	"testing"
)

var datePattern = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

// parseDate splits a date written as 2026-11-06 into its year, month and day.
func parseDate(s string) (year, month, day string, ok bool) {
	match := datePattern.FindStringSubmatch(s)
	if match == nil {
		return "", "", "", false
	}
	return match[1], match[2], match[3], true
}

func TestParseDate(t *testing.T) {
	year, month, day, ok := parseDate("due on 2026-11-06, at the latest")

	if !ok || year != "2026" || month != "11" || day != "06" {
		t.Errorf("parseDate = %q, %q, %q, %v, want 2026, 11, 06, true", year, month, day, ok)
	}
}

func TestParseDateWithoutADate(t *testing.T) {
	if _, _, _, ok := parseDate("no date here"); ok {
		t.Error("parseDate found a date in a text without one")
	}
}
