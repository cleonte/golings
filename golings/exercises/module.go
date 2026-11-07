package exercises

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// ErrNeedsInit means a module exercise was asked to run where its files are
// still stored with the .txt that keeps them out of the module of the
// repository: a clone, instead of the folder `golings init` makes.
var ErrNeedsInit = errors.New("module exercises run in the folder made by `golings init`")

// moduleDir is the folder of the exercise, as a path that works from any
// folder: the build and the run happen inside it.
func (e Exercise) moduleDir() string {
	dir := filepath.Join(e.Root, e.Dir())
	if abs, err := filepath.Abs(dir); err == nil {
		return abs
	}
	return dir
}

// checkModuleFolder fails with ErrNeedsInit when the folder has no go.mod or
// go.work, which is how a folder looks that was not made by `golings init`.
func (e Exercise) checkModuleFolder() error {
	for _, name := range []string{"go.mod", "go.work"} {
		if _, err := os.Stat(filepath.Join(e.moduleDir(), name)); err == nil {
			return nil
		}
	}
	return fmt.Errorf("%w: %s has no go.mod yet", ErrNeedsInit, e.Name)
}

// moduleEnv is the environment of the Go tool for a module exercise. Nothing
// may come from the network or from a Go that has to be downloaded, so the
// exercise works the same everywhere: a dependency has to be a folder on disk,
// and a go.mod that asks for a newer Go is an error. The workspace is off,
// except for an exercise that has a go.work of its own.
func (e Exercise) moduleEnv() []string {
	work := "off"
	if path := filepath.Join(e.moduleDir(), "go.work"); fileExists(path) {
		work = path
	}
	return []string{"GOWORK=" + work, "GOPROXY=off", "GOFLAGS=-mod=readonly", "GOTOOLCHAIN=local"}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
