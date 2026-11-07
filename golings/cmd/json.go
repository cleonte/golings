package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"path/filepath"

	"github.com/cleonte/golings/golings/exercises"
)

// JSONVersion is the version of the JSON the commands print with --json. It
// changes only when a field is renamed, removed or changes its meaning; a new
// field does not change it. The format is described in docs/json.md.
const JSONVersion = 1

// writeJSON prints v as one indented JSON document.
func writeJSON(w io.Writer, v any) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

// jsonError prints {"version":1,"error":"..."} for a command that could not
// give an answer (the info file is missing, the exercise does not exist) and
// returns the error that made the command exit non-zero.
func jsonError(w io.Writer, err error) error {
	if werr := writeJSON(w, struct {
		Version int    `json:"version"`
		Error   string `json:"error"`
	}{JSONVersion, err.Error()}); werr != nil {
		return werr
	}
	return err
}

// jsonFlagHelp is the help of the --json flag.
const jsonFlagHelp = "print the result as JSON, for editors and scripts (see docs/json.md)"

// ListJSON is what `golings list --json` prints.
type ListJSON struct {
	Version   int            `json:"version"`
	Done      int            `json:"done"`
	Total     int            `json:"total"`
	Topics    []TopicJSON    `json:"topics"`
	Exercises []ExerciseJSON `json:"exercises"`
}

// TopicJSON is the progress of one topic. Like Done and Total of the list, it
// leaves out the exercises this machine can't run.
type TopicJSON struct {
	Name  string `json:"name"`
	Done  int    `json:"done"`
	Total int    `json:"total"`
}

// ExerciseJSON describes one exercise of the list.
type ExerciseJSON struct {
	Name  string `json:"name"`
	Topic string `json:"topic"`
	// Path is the main file, relative to the folder that holds info.toml.
	Path string `json:"path"`
	Mode string `json:"mode"`
	// State is "pending" or "done".
	State string `json:"state"`
	// Skipped says why this machine can't run the exercise. It is left out when
	// it can; a skipped exercise is not counted in Done and Total.
	Skipped string `json:"skipped,omitempty"`
}

func exerciseJSON(e exercises.Exercise) ExerciseJSON {
	state := "pending"
	if e.State() == exercises.Done {
		state = "done"
	}
	return ExerciseJSON{
		Name:    e.Name,
		Topic:   e.Topic(),
		Path:    filepath.ToSlash(e.Path),
		Mode:    e.Mode,
		State:   state,
		Skipped: e.SkipReason(),
	}
}

// listJSON gathers everything `list --json` prints.
func listJSON(infoFile string) (ListJSON, error) {
	all, err := exercises.List(infoFile)
	if err != nil {
		return ListJSON{}, err
	}
	topics, err := exercises.ProgressByTopic(infoFile)
	if err != nil {
		return ListJSON{}, err
	}

	// Empty lists are [], not null, so a reader can loop over them.
	out := ListJSON{Version: JSONVersion, Topics: []TopicJSON{}, Exercises: []ExerciseJSON{}}
	for _, t := range topics {
		out.Done += t.Done
		out.Total += t.Total
		out.Topics = append(out.Topics, TopicJSON{Name: t.Topic, Done: t.Done, Total: t.Total})
	}
	for _, e := range all {
		out.Exercises = append(out.Exercises, exerciseJSON(e))
	}
	return out, nil
}

// RunJSON is what `golings run --json` prints.
type RunJSON struct {
	Version  int          `json:"version"`
	Exercise ExerciseJSON `json:"exercise"`
	// Status is one of:
	//   "passed"  the exercise built and ran, and it is marked done
	//   "pending" it passed, but still has the "I AM NOT DONE" marker
	//   "failed"  it did not build, its tests failed, or it crashed
	//   "timeout" it was stopped for running too long
	//   "skipped" this machine can't run it (an older Go, no race detector)
	Status string `json:"status"`
	// Output is what the program printed (stdout), for a pass.
	Output  string       `json:"output,omitempty"`
	Failure *FailureJSON `json:"failure,omitempty"`
	// Message is why the exercise was skipped.
	Message string `json:"message,omitempty"`
}

// FailureJSON explains a failed run.
type FailureJSON struct {
	// Kind is "compile", "test", "runtime" or "timeout".
	Kind string `json:"kind"`
	// Diagnostics are the compiler messages of a "compile" failure.
	Diagnostics []DiagnosticJSON `json:"diagnostics,omitempty"`
	// Output is the rest worth showing, without toolchain noise. Long panics and
	// data race reports are shortened; Full has all of it when they were.
	Output string `json:"output,omitempty"`
	Full   string `json:"full,omitempty"`
	// TimeoutSeconds is how long a "timeout" exercise was given.
	TimeoutSeconds float64 `json:"timeout_seconds,omitempty"`
}

// DiagnosticJSON is one compiler message.
type DiagnosticJSON struct {
	File    string `json:"file"`
	Line    int    `json:"line"`
	Column  int    `json:"column,omitempty"`
	Message string `json:"message"`
}

func failureJSON(f exercises.Failure) *FailureJSON {
	out := &FailureJSON{Output: f.Output, Full: f.Full, TimeoutSeconds: f.Timeout.Seconds()}
	switch f.Kind {
	case exercises.CompileFailure:
		out.Kind = "compile"
	case exercises.TestFailure:
		out.Kind = "test"
	case exercises.TimeoutFailure:
		out.Kind = "timeout"
	default:
		out.Kind = "runtime"
	}
	for _, d := range f.Diagnostics {
		out.Diagnostics = append(out.Diagnostics, DiagnosticJSON{File: filepath.ToSlash(d.File), Line: d.Line, Column: d.Col, Message: d.Message})
	}
	return out
}

// runJSON runs the exercise and describes what happened. A failing exercise is
// an answer, not an error: err is only set when the run could not be done.
func runJSON(e exercises.Exercise) (RunJSON, error) {
	out := RunJSON{Version: JSONVersion, Exercise: exerciseJSON(e)}

	result, err := e.Run()
	switch {
	case exercises.IsSkipped(err):
		out.Status, out.Message = "skipped", err.Error()
	case errors.Is(err, exercises.ErrNeedsInit):
		return out, err
	case err == nil && e.State() == exercises.Pending:
		out.Status, out.Output = "pending", result.Out
	case err == nil:
		out.Status, out.Output = "passed", result.Out
	case result.TimedOut:
		out.Status, out.Failure = "timeout", failureJSON(result.Failure())
	case result.BuildFailed || isExit(err):
		out.Status, out.Failure = "failed", failureJSON(result.Failure())
	default:
		return out, err
	}
	return out, nil
}

// isExit reports whether err is the exit status of a program that ran, as
// opposed to a run that could not be done (no temporary folder, cancelled).
func isExit(err error) bool {
	var exit *exec.ExitError
	return errors.As(err, &exit)
}

// runAsJSON prints what happened when running the exercise found (or not, with
// findErr) as JSON. It exits non-zero only when it can't say: the exercise does
// not exist, or it could not be run. A failing exercise is a "failed" status.
func runAsJSON(w io.Writer, e exercises.Exercise, name string, findErr error) error {
	if errors.Is(findErr, exercises.ErrExerciseNotFound) {
		return jsonError(w, fmt.Errorf("%w: %s", findErr, name))
	}
	if findErr != nil {
		return jsonError(w, findErr)
	}
	out, err := runJSON(e)
	if err != nil {
		return jsonError(w, err)
	}
	return writeJSON(w, out)
}
