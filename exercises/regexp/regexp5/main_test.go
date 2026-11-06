// regexp5
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"regexp"
	"testing"
)

// containsLiteral reports whether text contains needle exactly as it is written.
func containsLiteral(text, needle string) bool {
	// TODO: needle is used as a pattern, so "1+1" means "one or more 1, then 1" and "(" is not even a valid pattern. Turn the text into a pattern that matches itself.
	re, err := regexp.Compile(needle)
	if err != nil {
		return false
	}
	return re.MatchString(text)
}

func TestContainsLiteral(t *testing.T) {
	for _, tc := range []struct {
		text, needle string
		want         bool
	}{
		{"1+1=2", "1+1", true},
		{"axb", "a.b", false},
		{"a.b", "a.b", true},
		{"f(x)", "(", true},
	} {
		if got := containsLiteral(tc.text, tc.needle); got != tc.want {
			t.Errorf("containsLiteral(%q, %q) = %v, want %v", tc.text, tc.needle, got, tc.want)
		}
	}
}
