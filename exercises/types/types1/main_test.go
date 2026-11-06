// types1
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"testing"
)

// Direction is a kind of enum: a type with a few named values.
type Direction int

// TODO: The four values are written out and all are 0. Let iota count them 0, 1, 2, 3 instead: it goes up by one on each line of a const block, and a line without a value repeats the line above.
const (
	North Direction = 0
	East  Direction = 0
	South Direction = 0
	West  Direction = 0
)

func TestDirectionsCountUp(t *testing.T) {
	for want, direction := range []Direction{North, East, South, West} {
		if int(direction) != want {
			t.Errorf("the direction number %d has the value %d", want, direction)
		}
	}
}
