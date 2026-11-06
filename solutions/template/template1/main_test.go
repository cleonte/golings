// template1
// Make the tests pass!

package main_test

import (
	"strings"
	"testing"
	"text/template"
)

type greeting struct{ Name string }

// render fills in the greeting template with the name.
func render(name string) (string, error) {
	tmpl, err := template.New("greeting").Parse("Hello, {{.Name}}!")
	if err != nil {
		return "", err
	}

	var out strings.Builder
	if err := tmpl.Execute(&out, greeting{Name: name}); err != nil {
		return "", err
	}
	return out.String(), nil
}

func TestRender(t *testing.T) {
	got, err := render("Ada")

	if err != nil || got != "Hello, Ada!" {
		t.Errorf("render = %q, %v, want %q and no error", got, err, "Hello, Ada!")
	}
}
