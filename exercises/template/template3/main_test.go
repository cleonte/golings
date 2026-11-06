// template3
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"strings"
	"testing"
	"text/template"
)

// shout writes the text in capitals, followed by an exclamation mark.
func shout(text string) (string, error) {
	// TODO: The template calls a function named upper, but a template only knows its own functions and a few built-in ones. Register upper (strings.ToUpper will do) before parsing.
	tmpl, err := template.New("shout").Parse("{{upper .}}!")
	if err != nil {
		return "", err
	}

	var out strings.Builder
	if err := tmpl.Execute(&out, text); err != nil {
		return "", err
	}
	return out.String(), nil
}

func TestShout(t *testing.T) {
	got, err := shout("go")

	if err != nil || got != "GO!" {
		t.Errorf("shout = %q, %v, want %q and no error", got, err, "GO!")
	}
}
