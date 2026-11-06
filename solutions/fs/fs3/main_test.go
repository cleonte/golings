// fs3
// Make the tests pass!

package main_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

// totalSize adds up the size of every file below root, leaving out the .git folder.
func totalSize(root string) (int64, error) {
	var total int64

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && d.Name() == ".git" {
			return fs.SkipDir
		}
		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return err
			}
			total += info.Size()
		}
		return nil
	})
	return total, err
}

func TestTotalSizeLeavesOutGit(t *testing.T) {
	root := t.TempDir()
	files := map[string]int{
		"a.txt":                5,
		"sub/b.txt":            7,
		".git/objects/pack.db": 1000,
	}
	for name, size := range files {
		path := filepath.Join(root, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, make([]byte, size), 0o644); err != nil {
			t.Fatal(err)
		}
	}

	got, err := totalSize(root)

	if err != nil || got != 12 {
		t.Errorf("totalSize = %d, %v, want 12 and no error", got, err)
	}
}
