// Package doctor diagnoses the learner's setup: Go toolchain, workspace and
// the optional requirements of individual exercises.
package doctor

import (
	"fmt"
	"go/version"
	"os"
	"runtime"
	"strings"

	"github.com/cleonte/golings/golings/exercises"
)

// MinGoVersion is the oldest Go release golings supports.
const MinGoVersion = "1.26"

type Status int

const (
	OK Status = iota
	Warn
	Fail
)

type Check struct {
	Name   string
	Status Status
	Detail string
}

// Env holds everything doctor learns about the machine, so tests can fake it.
type Env struct {
	GOOS         string
	GoVersion    func() (string, error)
	CGOEnabled   func() bool
	HasCCompiler func() bool
	FileExists   func(path string) bool
}

func DefaultEnv() Env {
	race := exercises.DefaultRaceEnv()
	return Env{
		GOOS:         race.GOOS,
		GoVersion:    exercises.LocalGoVersion,
		CGOEnabled:   race.CGOEnabled,
		HasCCompiler: race.HasCCompiler,
		FileExists: func(path string) bool {
			_, err := os.Stat(path)
			return err == nil
		},
	}
}

// Run executes all checks and returns their results in display order.
func Run(infoFile string, env Env) []Check {
	checks := []Check{}

	local, err := env.GoVersion()
	checks = append(checks, goCheck(local, err))

	all, err := exercises.List(infoFile)
	if err != nil {
		return append(checks, Check{
			Name:   "Exercise manifest",
			Status: Fail,
			Detail: fmt.Sprintf("can't read %s: %v (run golings from the folder that holds it, or create one with `golings init`)", infoFile, err),
		})
	}
	checks = append(checks, Check{Name: "Exercise manifest", Status: OK, Detail: fmt.Sprintf("%s lists %d exercises", infoFile, len(all))})

	checks = append(checks, filesCheck(all, env))
	if local != "" {
		checks = append(checks, versionGatedCheck(all, local))
	}
	checks = append(checks, raceCheck(all, env))

	return checks
}

// Failed reports whether any check failed (warnings don't count).
func Failed(checks []Check) bool {
	for _, c := range checks {
		if c.Status == Fail {
			return true
		}
	}
	return false
}

func goCheck(local string, err error) Check {
	c := Check{Name: "Go toolchain"}
	switch {
	case err != nil:
		c.Status, c.Detail = Fail, fmt.Sprintf("can't run `go` (%v): is Go installed and on your PATH?", err)
	case version.Compare(local, "go"+MinGoVersion) < 0:
		c.Status, c.Detail = Fail, fmt.Sprintf("%s is too old, golings needs Go %s or newer", local, MinGoVersion)
	default:
		c.Status, c.Detail = OK, fmt.Sprintf("%s on %s/%s", local, runtime.GOOS, runtime.GOARCH)
	}
	return c
}

func filesCheck(all []exercises.Exercise, env Env) Check {
	var missing []string
	for _, ex := range all {
		if !env.FileExists(ex.Path) {
			missing = append(missing, ex.Name)
		}
	}
	if len(missing) > 0 {
		return Check{Name: "Exercise files", Status: Fail, Detail: fmt.Sprintf("%d missing: %s", len(missing), strings.Join(missing, ", "))}
	}
	return Check{Name: "Exercise files", Status: OK, Detail: "all present"}
}

func versionGatedCheck(all []exercises.Exercise, local string) Check {
	var tooNew []string
	for _, ex := range all {
		if !ex.Supports(local) {
			tooNew = append(tooNew, fmt.Sprintf("%s (needs Go %s)", ex.Name, ex.GoVersion))
		}
	}
	if len(tooNew) > 0 {
		return Check{Name: "Exercise Go versions", Status: Warn, Detail: fmt.Sprintf("%d will be skipped: %s", len(tooNew), strings.Join(tooNew, ", "))}
	}
	return Check{Name: "Exercise Go versions", Status: OK, Detail: "every exercise can run on this Go"}
}

// raceCheck looks at what `go test -race` needs. Only the exercises marked
// with race = true are built with it; the others run anywhere Go runs.
func raceCheck(all []exercises.Exercise, env Env) Check {
	c := Check{Name: "Race detector"}
	var needing []string
	for _, ex := range all {
		if ex.Race {
			needing = append(needing, ex.Name)
		}
	}

	race := exercises.RaceEnv{GOOS: env.GOOS, CGOEnabled: env.CGOEnabled, HasCCompiler: env.HasCCompiler}
	switch {
	case len(needing) == 0:
		c.Status, c.Detail = OK, "not needed"
	case race.Supported():
		c.Status, c.Detail = OK, fmt.Sprintf("available, for the %d exercises that need it", len(needing))
	default:
		c.Status = Warn
		c.Detail = fmt.Sprintf("not available (it needs cgo and a C compiler: install gcc or clang and set CGO_ENABLED=1); %d will be skipped: %s", len(needing), strings.Join(needing, ", "))
	}
	return c
}
