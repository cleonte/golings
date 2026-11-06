// Package workspace creates a folder with the exercises and restores single
// exercises to their original state, from the copy embedded in the binary.
package workspace

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/cleonte/golings/golings/doctor"
	"github.com/cleonte/golings/golings/exercises"
)

var ErrExists = errors.New("already exists")

// Init writes the exercises in src to a new folder dest, with a go.mod so
// that the Go tools and editors treat it as a module. It refuses to touch a
// path that already exists.
func Init(src fs.FS, dest string) error {
	if _, err := os.Stat(dest); err == nil {
		return fmt.Errorf("%s %w", dest, ErrExists)
	} else if !errors.Is(err, fs.ErrNotExist) {
		return err
	}

	if err := write(src, dest); err != nil {
		_ = os.RemoveAll(dest)
		return err
	}
	return nil
}

func write(src fs.FS, dest string) error {
	err := fs.WalkDir(src, ".", func(name string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Names inside an fs.FS always use forward slashes.
		target := filepath.Join(dest, filepath.FromSlash(name))
		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		data, err := fs.ReadFile(src, name)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, 0o644)
	})
	if err != nil {
		return err
	}

	goMod := fmt.Sprintf("module %s\n\ngo %s\n", doctor.ExercisesModule, doctor.MinGoVersion)
	return os.WriteFile(filepath.Join(dest, "go.mod"), []byte(goMod), 0o644)
}

// Reset restores the exercise called name below root to the version in src: all
// the files of its folder come back, the Go files added to the folder go, and
// the "I AM NOT DONE" marker is back.
// It returns the path of the main file, relative to root and with forward
// slashes.
func Reset(src fs.FS, root, name string) (string, error) {
	info, err := fs.ReadFile(src, "info.toml")
	if err != nil {
		return "", err
	}
	all, err := exercises.ParseInfo(info)
	if err != nil {
		return "", err
	}

	for _, ex := range all {
		if ex.Name != name {
			continue
		}
		// The path comes from the manifest, never from the user's input.
		if !fs.ValidPath(ex.Path) {
			return "", fmt.Errorf("exercise %s has an invalid path %q", ex.Name, ex.Path)
		}
		// The folder is the exercise: every file in it comes back, and the Go
		// files the learner added to it go, so the exercise starts over.
		dir := path.Dir(ex.Path)
		restored := map[string]bool{}
		err := fs.WalkDir(src, dir, func(name string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			target := filepath.Join(root, filepath.FromSlash(name))
			if d.IsDir() {
				return os.MkdirAll(target, 0o755)
			}
			original, err := fs.ReadFile(src, name)
			if err != nil {
				return err
			}
			restored[target] = true
			return os.WriteFile(target, original, 0o644)
		})
		if err != nil {
			return ex.Path, err
		}
		err = filepath.WalkDir(filepath.Join(root, filepath.FromSlash(dir)), func(name string, d fs.DirEntry, err error) error {
			if err != nil || d.IsDir() || restored[name] || filepath.Ext(name) != ".go" {
				return err
			}
			return os.Remove(name)
		})
		return ex.Path, err
	}

	return "", exercises.ErrExerciseNotFound
}
