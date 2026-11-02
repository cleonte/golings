package exercises

import (
	"errors"
	"os"

	"github.com/pelletier/go-toml/v2"
)

var ErrExerciseNotFound = errors.New("exercise not found")
var ErrNoPendingExercises = errors.New("no pending exercises")

type Info struct {
	Exercises []Exercise
}

func List(infoFile string) ([]Exercise, error) {
	data, err := os.ReadFile(infoFile)
	if err != nil {
		return nil, err
	}

	return ParseInfo(data)
}

// ParseInfo reads the exercises from the contents of an info.toml.
func ParseInfo(data []byte) ([]Exercise, error) {
	var info Info
	if err := toml.Unmarshal(data, &info); err != nil {
		return info.Exercises, err
	}

	return info.Exercises, nil
}

// runnable reports whether this machine can run the exercise: the local Go is
// new enough (if it can't be determined, it is assumed to be) and the race
// detector is there when the exercise needs it.
func runnable(e Exercise) bool {
	return e.SkipReason() == ""
}

// Skipped lists the exercises this machine can't run: the local Go is too old
// or the race detector is missing.
func Skipped(infoFile string) ([]Exercise, error) {
	allExercises, err := List(infoFile)
	if err != nil {
		return nil, err
	}

	var skipped []Exercise
	for _, exercise := range allExercises {
		if !runnable(exercise) {
			skipped = append(skipped, exercise)
		}
	}
	return skipped, nil
}

func NextPending(infoFile string) (Exercise, error) {
	allExercises, err := List(infoFile)
	if err != nil {
		return Exercise{}, err
	}

	for _, exercise := range allExercises {
		if exercise.State() == Pending && runnable(exercise) {
			return exercise, nil
		}
	}

	return Exercise{}, ErrNoPendingExercises
}

func Find(exercise string, infoFile string) (Exercise, error) {
	exs, err := List(infoFile)
	if err != nil {
		return Exercise{}, err
	}

	for _, ex := range exs {
		if ex.Name == exercise {
			return ex, nil
		}
	}

	return Exercise{}, ErrExerciseNotFound
}

func Progress(infoFile string) (float32, int, int, error) {
	allExercises, err := List(infoFile)
	if err != nil {
		return 0.0, 0, 0, err
	}
	totalDone, total := 0, 0
	for _, exercise := range allExercises {
		if !runnable(exercise) {
			continue
		}
		total++
		if exercise.State() == Done {
			totalDone++
		}
	}

	if total == 0 {
		return 0, 0, 0, nil
	}

	return float32(totalDone) / float32(total), totalDone, total, nil
}
