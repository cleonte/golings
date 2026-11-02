package exercises

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

// ErrRaceUnsupported means the exercise needs the race detector and this
// machine can't provide it.
var ErrRaceUnsupported = errors.New("the race detector is not available here")

// IsSkipped reports whether err says an exercise was not run because the
// machine can't run it: Go too old, or no race detector.
func IsSkipped(err error) bool {
	return errors.Is(err, ErrGoVersionTooOld) || errors.Is(err, ErrRaceUnsupported)
}

// RaceEnv holds what decides whether `go test -race` works, so tests (and
// doctor) can fake the machine.
type RaceEnv struct {
	GOOS         string
	CGOEnabled   func() bool
	HasCCompiler func() bool
}

// DefaultRaceEnv looks at this machine.
func DefaultRaceEnv() RaceEnv {
	return RaceEnv{
		GOOS: runtime.GOOS,
		CGOEnabled: func() bool {
			out, err := exec.Command("go", "env", "CGO_ENABLED").Output()
			return err == nil && strings.TrimSpace(string(out)) == "1"
		},
		HasCCompiler: func() bool {
			// The compiler cgo would use, then the usual names.
			candidates := []string{"cc", "gcc", "clang"}
			if out, err := exec.Command("go", "env", "CC").Output(); err == nil {
				if fields := strings.Fields(string(out)); len(fields) > 0 {
					candidates = append([]string{fields[0]}, candidates...)
				}
			}
			for _, cc := range candidates {
				if _, err := exec.LookPath(cc); err == nil {
					return true
				}
			}
			return false
		},
	}
}

// Supported reports whether `go test -race` will work: on macOS it always does
// (it doesn't need cgo there), elsewhere it needs cgo and a C compiler.
func (r RaceEnv) Supported() bool {
	if r.GOOS == "darwin" {
		return true
	}
	return r.CGOEnabled() && r.HasCCompiler()
}

// RaceSupported answers for this machine, looked up once per process.
var RaceSupported = sync.OnceValue(func() bool { return DefaultRaceEnv().Supported() })

// SkipReason says why the exercise can't run here, or is empty if it can.
func (e Exercise) SkipReason() string {
	if local, err := LocalGoVersion(); err == nil && !e.Supports(local) {
		return "needs Go " + e.GoVersion
	}
	if e.Race && !RaceSupported() {
		return "needs the race detector"
	}
	return ""
}

// skipError is the error RunContext gives for an exercise that can't run here.
func (e Exercise) skipError() error {
	if local, err := LocalGoVersion(); err == nil && !e.Supports(local) {
		return fmt.Errorf("%w: %s needs Go %s, found %s", ErrGoVersionTooOld, e.Name, e.GoVersion, local)
	}
	if e.Race && !RaceSupported() {
		return fmt.Errorf("%w: %s needs it. Install a C compiler (gcc or clang) and set CGO_ENABLED=1, then try `golings doctor`", ErrRaceUnsupported, e.Name)
	}
	return nil
}
