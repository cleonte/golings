package exercises

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var notDoneRegex = regexp.MustCompile(`(?m)^\s*///?\s*I\s+AM\s+NOT\s+DONE`)

type Exercise struct {
	Name string
	Path string
	Mode string
	Hint string
	// Nudge and Stronger are optional, smaller hints that come before Hint: what
	// the failure means, and where to change. Stronger is only set with Nudge.
	Nudge    string
	Stronger string
	// GoVersion is the minimum Go version (e.g. "1.26") needed to run the exercise.
	GoVersion string `toml:"go_version"`
	// Race is set for the exercises that only fail with the race detector. They
	// are built with -race, and skipped where that is not available.
	Race bool `toml:"race"`
	// Module is set for the exercises about modules: the folder of the exercise
	// is a module of its own (it has a go.mod, maybe a go.work and modules
	// inside), and the exercise is built and run in it. The files of such a
	// folder are stored with a .txt added to their names (go.mod.txt), in the
	// repository and in the binary, so that they are not taken for part of this
	// module or for a module of their own; `golings init` and `golings reset`
	// write them without it, and Path is the name without it.
	Module bool `toml:"module"`
	// Root is the folder Path is relative to, when it is not the current one.
	// It is never read from the manifest: tests that build an exercise in a
	// folder of their own set it.
	Root string `toml:"-"`
}

// Dir is the folder of the exercise, relative to Root. The whole folder is the
// exercise: Path names its main file, the one with the "I AM NOT DONE" marker,
// and any other Go file in the folder is part of the same package.
func (e Exercise) Dir() string {
	return filepath.Dir(e.Path)
}

// storedSuffix is what the files of a module exercise have added to their names
// in the repository and in the binary.
const storedSuffix = ".txt"

// StoredName is the name of a file of the exercise folder where the exercises
// are stored (the repository, the binary): the same name, with .txt added for a
// module exercise.
func (e Exercise) StoredName(name string) string {
	if e.Module {
		return name + storedSuffix
	}
	return name
}

// LearnerName is the name a stored file has in the folder of a learner: for a
// module exercise without the .txt, as it is in the manifest.
func (e Exercise) LearnerName(name string) string {
	if e.Module {
		return strings.TrimSuffix(name, storedSuffix)
	}
	return name
}

// HintLevels returns the hints of the exercise from the smallest to the full
// one: the nudge and the stronger hint when there are any, then the hint.
func (e Exercise) HintLevels() []string {
	var levels []string
	for _, hint := range []string{e.Nudge, e.Stronger, e.Hint} {
		if hint = strings.TrimSpace(hint); hint != "" {
			levels = append(levels, hint)
		}
	}
	return levels
}

// ErrNotMarked means the exercise file has no "I AM NOT DONE" marker to remove.
var ErrNotMarked = errors.New("the exercise has no 'I AM NOT DONE' marker")

// MarkDone removes the "I AM NOT DONE" marker from the exercise file, which is
// what a learner does by hand to say that they have finished an exercise.
func (e Exercise) MarkDone() error {
	info, err := os.Stat(e.Path)
	if err != nil {
		return err
	}
	data, err := os.ReadFile(e.Path)
	if err != nil {
		return err
	}

	lines := strings.SplitAfter(string(data), "\n")
	kept := make([]string, 0, len(lines))
	removed := false
	for _, line := range lines {
		if notDoneRegex.MatchString(line) {
			removed = true
			continue
		}
		kept = append(kept, line)
	}
	if !removed {
		return ErrNotMarked
	}

	return os.WriteFile(e.Path, []byte(strings.Join(kept, "")), info.Mode().Perm())
}

func (e Exercise) State() State {
	data, err := os.ReadFile(e.Path)
	if err != nil {
		return Pending
	}

	if notDoneRegex.Match(data) {
		return Pending
	}
	return Done
}

type State int

const (
	Pending State = iota + 1
	Done
)

func (s State) String() string {
	return [...]string{"Pending", "Done"}[s-1]
}
