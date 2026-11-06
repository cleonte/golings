// fs4
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"os"
	"path/filepath"
	"testing"
)

// readConfig reads the file called name inside dir. name comes from outside, so it must not lead out of dir.
func readConfig(dir, name string) ([]byte, error) {
	// TODO: A name such as ../secret.txt leaves dir. Open dir as an os.Root and read through it: it refuses a path that escapes.
	return os.ReadFile(filepath.Join(dir, name))
}

func setup(t *testing.T) (public string) {
	t.Helper()
	base := t.TempDir()
	public = filepath.Join(base, "public")
	if err := os.MkdirAll(public, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(public, "config.txt"), []byte("ok"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(base, "secret.txt"), []byte("secret"), 0o644); err != nil {
		t.Fatal(err)
	}
	return public
}

func TestAFileInsideTheFolderIsRead(t *testing.T) {
	got, err := readConfig(setup(t), "config.txt")

	if err != nil || string(got) != "ok" {
		t.Errorf("readConfig = %q, %v, want %q and no error", got, err, "ok")
	}
}

func TestAPathOutOfTheFolderIsRefused(t *testing.T) {
	got, err := readConfig(setup(t), "../secret.txt")

	if err == nil {
		t.Errorf("readConfig returned %q for a file outside the folder, want an error", got)
	}
}
