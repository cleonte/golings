// template3
// Make the tests pass!

package main_test

import (
	"strings"
	"testing"
	"text/template"
)

// shout writes the text in capitals, followed by an exclamation mark.
func shout(text string) (string, error) {
	tmpl, err := template.New("shout").Funcs(template.FuncMap{"upper": strings.ToUpper}).Parse("{{upper .}}!")
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
