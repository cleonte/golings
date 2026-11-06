// fs1
// Make the tests pass!

package main_test

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

// readOrDefault returns the content of the file at path, or def when there is no such file. Any other problem is an error.
func readOrDefault(path, def string) (string, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, fs.ErrNotExist) {
		return def, nil
	}
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func TestAMissingFileGivesTheDefault(t *testing.T) {
	got, err := readOrDefault(filepath.Join(t.TempDir(), "missing.txt"), "fallback")

	if err != nil || got != "fallback" {
		t.Errorf("readOrDefault = %q, %v, want %q and no error", got, err, "fallback")
	}
}

func TestAnExistingFileGivesItsContent(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.txt")
	if err := os.WriteFile(path, []byte("from the file"), 0o644); err != nil {
		t.Fatal(err)
	}

	got, err := readOrDefault(path, "fallback")

	if err != nil || got != "from the file" {
		t.Errorf("readOrDefault = %q, %v, want %q and no error", got, err, "from the file")
	}
}

func TestOtherErrorsAreNotHidden(t *testing.T) {
	// A folder exists, but it can't be read as a file.
	if _, err := readOrDefault(t.TempDir(), "fallback"); err == nil {
		t.Error("readOrDefault returned no error for a folder: only a missing file should give the default")
	}
}
