// template5
// Make the tests pass!

package main_test

import (
	"html/template"
	"strings"
	"testing"
)

// page puts a comment written by a visitor in a paragraph of a web page.
func page(comment string) (string, error) {
	tmpl, err := template.New("page").Parse("<p>{{.}}</p>")
	if err != nil {
		return "", err
	}

	var out strings.Builder
	if err := tmpl.Execute(&out, comment); err != nil {
		return "", err
	}
	return out.String(), nil
}

func TestPageEscapesTheComment(t *testing.T) {
	got, err := page("<script>alert(1)</script>")

	want := "<p>&lt;script&gt;alert(1)&lt;/script&gt;</p>"
	if err != nil || got != want {
		t.Errorf("page = %q, %v, want %q and no error", got, err, want)
	}
}
