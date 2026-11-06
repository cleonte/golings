// regexp1
// Make the tests pass!

package main_test

import (
	"regexp"
	"testing"
)

// isZip reports whether s is exactly a five-digit zip code.
func isZip(s string) bool {
	return regexp.MustCompile(`^[0-9]{5}$`).MatchString(s)
}

func TestIsZip(t *testing.T) {
	for _, tc := range []struct {
		text string
		want bool
	}{
		{"12345", true},
		{"1234", false},
		{"a12345b", false},
		{"123456", false},
	} {
		if got := isZip(tc.text); got != tc.want {
			t.Errorf("isZip(%q) = %v, want %v", tc.text, got, tc.want)
		}
	}
}
