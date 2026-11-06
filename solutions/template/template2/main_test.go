// template2
// Make the tests pass!

package main_test

import (
	"strings"
	"testing"
	"text/template"
)

type shopping struct{ Items []string }

// list writes one line, "- item", for each item.
func list(items []string) (string, error) {
	tmpl, err := template.New("list").Parse("{{range .Items}}- {{.}}\n{{end}}")
	if err != nil {
		return "", err
	}

	var out strings.Builder
	if err := tmpl.Execute(&out, shopping{Items: items}); err != nil {
		return "", err
	}
	return out.String(), nil
}

func TestList(t *testing.T) {
	got, err := list([]string{"milk", "eggs"})

	if err != nil || got != "- milk\n- eggs\n" {
		t.Errorf("list = %q, %v, want %q and no error", got, err, "- milk\n- eggs\n")
	}
}

func TestEmptyList(t *testing.T) {
	got, err := list(nil)

	if err != nil || got != "" {
		t.Errorf("list(nil) = %q, %v, want an empty string and no error", got, err)
	}
}
