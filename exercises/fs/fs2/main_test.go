// fs2
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"io/fs"
	"os"
	"strings"
	"testing"
	"testing/fstest"
)

// countGoFiles counts the .go files anywhere in fsys, in folders inside folders too.
func countGoFiles(fsys fs.FS) (int, error) {
	// TODO: This looks at the real current folder and ignores fsys. Walk fsys itself from ".", and count the files whose name ends in .go.
	entries, err := os.ReadDir(".")
	if err != nil {
		return 0, err
	}

	count := 0
	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".go") {
			count++
		}
	}
	return count, nil
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
