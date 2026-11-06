// fuzz1
// Make the tests pass!

package main_test

import (
	"slices"
	"testing"
	"unicode/utf8"
)

// Reverse returns s with its characters in the opposite order.
func Reverse(s string) string {
	runes := []rune(s)
	slices.Reverse(runes)
	return string(runes)
}

// FuzzReverse checks a property that must hold for any text. A normal
// "go test" runs it on the seeds below; "go test -fuzz FuzzReverse" would
// also try thousands of inputs it makes up.
func FuzzReverse(f *testing.F) {
	f.Add("hello")
	f.Add("héllo, 世界")
	f.Fuzz(func(t *testing.T, s string) {
		if !utf8.ValidString(s) {
			t.Skip()
		}
		rev := Reverse(s)
		if !utf8.ValidString(rev) {
			t.Errorf("Reverse(%q) = %q is not valid UTF-8", s, rev)
		}
		if Reverse(rev) != s {
			t.Errorf("reversing %q twice gave %q", s, Reverse(rev))
		}
	})
}
