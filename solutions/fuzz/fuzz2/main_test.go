// fuzz2
// Make the tests pass!

package main_test

import (
	"errors"
	"strings"
	"testing"
)

var errNoEquals = errors.New("no = in the text")

// ParseKV splits "key=value" into its two parts.
func ParseKV(s string) (key, value string, err error) {
	key, value, ok := strings.Cut(s, "=")
	if !ok {
		return "", "", errNoEquals
	}
	return key, value, nil
}

func FuzzParseKV(f *testing.F) {
	f.Add("name=gopher")
	f.Add("empty=")
	f.Add("novalue")
	f.Add("a=b=c")
	f.Fuzz(func(t *testing.T, s string) {
		key, value, err := ParseKV(s)
		if err != nil {
			return // an error is fine, a panic is not
		}
		if key+"="+value != s {
			t.Errorf("ParseKV(%q) = %q and %q: they do not make the text again", s, key, value)
		}
	})
}
