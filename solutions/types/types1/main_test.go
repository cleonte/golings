// types1
// Make the tests pass!

package main_test

import (
	"testing"
)

// Direction is a kind of enum: a type with a few named values.
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func TestDirectionsCountUp(t *testing.T) {
	for want, direction := range []Direction{North, East, South, West} {
		if int(direction) != want {
			t.Errorf("the direction number %d has the value %d", want, direction)
		}
	}
}
