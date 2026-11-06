// fs5
// Make the tests pass!

package main_test

import (
	"io/fs"
	"slices"
	"testing"
	"testing/fstest"
)

// cssFiles lists the style sheets in assets/css, by their name inside that folder (app.css, not assets/css/app.css).
func cssFiles(fsys fs.FS) ([]string, error) {
	sub, err := fs.Sub(fsys, "assets/css")
	if err != nil {
		return nil, err
	}
	return fs.Glob(sub, "*.css")
}

func TestCSSFiles(t *testing.T) {
	fsys := fstest.MapFS{
		"index.css":            {Data: []byte("body {}")},
		"assets/css/app.css":   {Data: []byte("a {}")},
		"assets/css/reset.css": {Data: []byte("* {}")},
		"assets/js/app.js":     {Data: []byte("1")},
	}

	got, err := cssFiles(fsys)

	want := []string{"app.css", "reset.css"}
	if err != nil || !slices.Equal(got, want) {
		t.Errorf("cssFiles = %q, %v, want %q and no error", got, err, want)
	}
}
