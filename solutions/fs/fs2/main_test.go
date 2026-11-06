// fs2
// Make the tests pass!

package main_test

import (
	"io/fs"
	"strings"
	"testing"
	"testing/fstest"
)

// countGoFiles counts the .go files anywhere in fsys, in folders inside folders too.
func countGoFiles(fsys fs.FS) (int, error) {
	count := 0
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".go") {
			count++
		}
		return nil
	})
	return count, err
}

func TestCountGoFiles(t *testing.T) {
	// A file system in memory: no files on the disk are needed.
	fsys := fstest.MapFS{
		"main.go":         {Data: []byte("package main")},
		"cmd/run.go":      {Data: []byte("package cmd")},
		"cmd/sub/deep.go": {Data: []byte("package sub")},
		"README.md":       {Data: []byte("# readme")},
		"cmd/notes.txt":   {Data: []byte("notes")},
	}

	got, err := countGoFiles(fsys)

	if err != nil || got != 3 {
		t.Errorf("countGoFiles = %d, %v, want 3 and no error", got, err)
	}
}
