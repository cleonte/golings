// regexp3
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"regexp"
	"testing"
)

var namePattern = regexp.MustCompile(`(\w+), (\w+)`)

// firstLast turns "Lovelace, Ada" into "Ada Lovelace".
func firstLast(s string) string {
	// TODO: This puts the text "${2} ${1}" in the result as it is. A replacement that expands the groups has a different method.
	return namePattern.ReplaceAllLiteralString(s, "${2} ${1}")
}

func TestFirstLast(t *testing.T) {
	got := firstLast("Lovelace, Ada and Turing, Alan")

	if want := "Ada Lovelace and Alan Turing"; got != want {
		t.Errorf("firstLast = %q, want %q", got, want)
	}
}
