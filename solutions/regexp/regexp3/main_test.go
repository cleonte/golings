// regexp3
// Make the tests pass!

package main_test

import (
	"regexp"
	"testing"
)

var namePattern = regexp.MustCompile(`(\w+), (\w+)`)

// firstLast turns "Lovelace, Ada" into "Ada Lovelace".
func firstLast(s string) string {
	return namePattern.ReplaceAllString(s, "${2} ${1}")
}

func TestFirstLast(t *testing.T) {
	got := firstLast("Lovelace, Ada and Turing, Alan")

	if want := "Ada Lovelace and Alan Turing"; got != want {
		t.Errorf("firstLast = %q, want %q", got, want)
	}
}
