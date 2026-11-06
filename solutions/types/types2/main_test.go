// types2
// Make the tests pass!

package main_test

import (
	"fmt"
	"testing"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

// String makes fmt print a Direction by its name.
func (d Direction) String() string {
	names := [...]string{"North", "East", "South", "West"}
	if d < 0 || int(d) >= len(names) {
		return fmt.Sprintf("Direction(%d)", int(d))
	}
	return names[d]
}

func TestDirectionsPrintTheirNames(t *testing.T) {
	for _, tc := range []struct {
		direction Direction
		want      string
	}{{North, "North"}, {East, "East"}, {South, "South"}, {West, "West"}} {
		if got := fmt.Sprint(tc.direction); got != tc.want {
			t.Errorf("fmt.Sprint(%d) = %q, want %q", int(tc.direction), got, tc.want)
		}
	}
}

func TestAnUnknownDirection(t *testing.T) {
	if got := fmt.Sprint(Direction(9)); got != "Direction(9)" {
		t.Errorf("fmt.Sprint(Direction(9)) = %q, want %q", got, "Direction(9)")
	}
}
