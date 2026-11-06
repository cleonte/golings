package exercises

import (
	"slices"
	"testing"
)

func TestBuildArgsUsesTheRaceDetectorOnlyWhenTheExerciseNeedsIt(t *testing.T) {
	tests := []struct {
		name string
		ex   Exercise
		want []string
	}{
		{"compile", Exercise{Path: "a/main.go", Mode: "compile"}, []string{"build", "-o", "out", "./a"}},
		{"compile ignores race", Exercise{Path: "a/main.go", Mode: "compile", Race: true}, []string{"build", "-o", "out", "./a"}},
		{"test", Exercise{Path: "a/main_test.go", Mode: "test"}, []string{"test", "-c", "-o", "out", "./a"}},
		{"test with race", Exercise{Path: "a/main_test.go", Mode: "test", Race: true}, []string{"test", "-c", "-race", "-o", "out", "./a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildArgs(tt.ex, "out"); !slices.Equal(got, tt.want) {
				t.Errorf("buildArgs = %v, want %v", got, tt.want)
			}
		})
	}
}
