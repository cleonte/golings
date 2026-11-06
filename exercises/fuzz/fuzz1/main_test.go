// fuzz1
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"testing"
	"unicode/utf8"
)

// Reverse returns s with its characters in the opposite order.
func Reverse(s string) string {
	// TODO: This reverses the bytes, and a character such as é or 世 takes several bytes, so reversing them breaks the character. Reverse the runes instead.
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
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
