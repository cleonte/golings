// template4
// Make the tests pass!

package main_test

import (
	"strings"
	"testing"
	"text/template"
)

// join writes the words separated by a comma and a space: "a, b, c".
func join(words []string) (string, error) {
	tmpl, err := template.New("join").Parse("{{range $i, $word := .}}{{if $i}}, {{end}}{{$word}}{{end}}")
	if err != nil {
		return "", err
	}

	var out strings.Builder
	if err := tmpl.Execute(&out, words); err != nil {
		return "", err
	}
	return out.String(), nil
}

func TestJoin(t *testing.T) {
	got, err := join([]string{"a", "b", "c"})

	if err != nil || got != "a, b, c" {
		t.Errorf("join = %q, %v, want %q and no error", got, err, "a, b, c")
	}
}

func TestJoinOfOneWord(t *testing.T) {
	got, err := join([]string{"only"})

	if err != nil || got != "only" {
		t.Errorf("join = %q, %v, want %q and no error", got, err, "only")
	}
}
