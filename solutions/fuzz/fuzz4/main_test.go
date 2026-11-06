// fuzz4
// Make the tests pass!

package main_test

import (
	"strings"
	"testing"
)

// Escape makes s safe to put on one line: a newline becomes \n and a
// backslash becomes \\.
func Escape(s string) string {
	return strings.NewReplacer(`\`, `\\`, "\n", `\n`).Replace(s)
}

// Unescape is the opposite of Escape.
func Unescape(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' && i+1 < len(s) {
			switch s[i+1] {
			case 'n':
				b.WriteByte('\n')
				i++
				continue
			case '\\':
				b.WriteByte('\\')
				i++
				continue
			}
		}
		b.WriteByte(s[i])
	}
	return b.String()
}

// FuzzEscapeRoundTrip checks that going there and back gives the original.
func FuzzEscapeRoundTrip(f *testing.F) {
	f.Add("plain")
	f.Add("two\nlines")
	f.Add(`back\slash`)
	f.Add(`a literal \n in the text`)
	f.Add(`ends with a backslash\`)
	f.Fuzz(func(t *testing.T, s string) {
		escaped := Escape(s)
		if strings.Contains(escaped, "\n") {
			t.Errorf("Escape(%q) = %q still has a newline", s, escaped)
		}
		if back := Unescape(escaped); back != s {
			t.Errorf("Unescape(Escape(%q)) = %q", s, back)
		}
	})
}
