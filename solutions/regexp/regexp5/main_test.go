// regexp5
// Make the tests pass!

package main_test

import (
	"regexp"
	"testing"
)

// containsLiteral reports whether text contains needle exactly as it is written.
func containsLiteral(text, needle string) bool {
	re, err := regexp.Compile(regexp.QuoteMeta(needle))
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
