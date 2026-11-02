// go127_1
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"strings"
	"testing"
)

// splitExtension splits a file name at its last dot: "archive.tar.gz" is
// "archive.tar" and "gz". A name without a dot has no extension.
func splitExtension(name string) (base, ext string) {
	// TODO: Cut splits at the first dot. Split at the last one instead (Go 1.27).
	base, ext, _ = strings.Cut(name, ".")
	return base, ext
}

// splitPath splits a path into its folder and its file name: "docs/go/intro.md"
// is "docs/go" and "intro.md". A path without a slash has no folder.
func splitPath(path string) (dir, file string) {
	// TODO: Split at the last slash. When there is none, the whole path is the file name.
	dir, file, _ = strings.Cut(path, "/")
	return dir, file
}

func TestSplitExtension(t *testing.T) {
	tests := []struct {
		name     string
		wantBase string
		wantExt  string
	}{
		{name: "notes.txt", wantBase: "notes", wantExt: "txt"},
		{name: "archive.tar.gz", wantBase: "archive.tar", wantExt: "gz"},
		{name: "README", wantBase: "README", wantExt: ""},
	}

	for _, tc := range tests {
		base, ext := splitExtension(tc.name)
		if base != tc.wantBase || ext != tc.wantExt {
			t.Errorf("splitExtension(%q) = %q, %q, want %q, %q", tc.name, base, ext, tc.wantBase, tc.wantExt)
		}
	}
}

func TestSplitPath(t *testing.T) {
	tests := []struct {
		path     string
		wantDir  string
		wantFile string
	}{
		{path: "docs/go/intro.md", wantDir: "docs/go", wantFile: "intro.md"},
		{path: "docs/intro.md", wantDir: "docs", wantFile: "intro.md"},
		{path: "intro.md", wantDir: "", wantFile: "intro.md"},
		{path: "/intro.md", wantDir: "", wantFile: "intro.md"},
	}

	for _, tc := range tests {
		dir, file := splitPath(tc.path)
		if dir != tc.wantDir || file != tc.wantFile {
			t.Errorf("splitPath(%q) = %q, %q, want %q, %q", tc.path, dir, file, tc.wantDir, tc.wantFile)
		}
	}
}
