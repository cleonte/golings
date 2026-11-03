package exercises_test

import (
	"errors"
	"os"
	"testing"

	"github.com/cleonte/golings/golings/exercises"
)

// TestWithoutTheRaceDetector checks what a learner on a machine without it gets
// (no C compiler, or CGO_ENABLED=0 on Linux or Windows) with the real manifest:
// the exercises with race = true are skipped, not failed. The test needs such a
// machine, so it only runs when GOLINGS_EXPECT_NO_RACE is set; the "learner
// environments" workflow sets it in a job with CGO_ENABLED=0.
func TestWithoutTheRaceDetector(t *testing.T) {
	if os.Getenv("GOLINGS_EXPECT_NO_RACE") == "" {
		t.Skip("set GOLINGS_EXPECT_NO_RACE=1 on a machine without the race detector (CGO_ENABLED=0) to run this")
	}
	if exercises.RaceSupported() {
		t.Fatal("this environment is meant to have no race detector (CGO_ENABLED=0), but it has one")
	}

	infoFile := "../../info.toml"
	all, err := exercises.List(infoFile)
	if err != nil {
		t.Fatal(err)
	}

	racy := 0
	for _, ex := range all {
		if !ex.Race {
			continue
		}
		racy++

		if reason := ex.SkipReason(); reason != "needs the race detector" {
			t.Errorf("%s: SkipReason() = %q, want it skipped for the race detector", ex.Name, reason)
		}
		result, err := ex.Run()
		if !errors.Is(err, exercises.ErrRaceUnsupported) || result.BuildFailed {
			t.Errorf("%s: Run() = (built: %t) %v, want a skip with ErrRaceUnsupported, and no build", ex.Name, result.BuildFailed, err)
		}
	}
	if racy == 0 {
		t.Fatal("no exercise in the manifest has race = true: this test checks nothing")
	}

	skipped, err := exercises.Skipped(infoFile)
	if err != nil {
		t.Fatal(err)
	}
	if len(skipped) != racy {
		t.Errorf("%d exercises skipped, want the %d with race = true", len(skipped), racy)
	}
	_, _, total, err := exercises.Progress(infoFile)
	if err != nil {
		t.Fatal(err)
	}
	if total != len(all)-racy {
		t.Errorf("the progress counts %d exercises, want %d (the skipped ones are left out)", total, len(all)-racy)
	}
}
