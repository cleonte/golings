package exercises

import (
	"errors"
	"os"
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
