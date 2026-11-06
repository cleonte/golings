// reflect1
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"reflect"
	"testing"
)

// kindOf returns the kind of a value: the shape of its type, such as "int",
// "slice" or "struct", whatever name the type was given.
func kindOf(v any) string {
	// TODO: The type name is not the kind: Celsius is a type name, float64 is its kind. Ask for the Kind of the value and return its text.
	return reflect.TypeOf(v).String()
}

type Celsius float64

type Point struct{ X, Y int }

func TestKindIgnoresTheNameOfTheType(t *testing.T) {
	number := 7
	for _, tc := range []struct {
		value any
		want  string
	}{
		{Celsius(21.5), "float64"},
		{Point{1, 2}, "struct"},
		{[]string{"a"}, "slice"},
		{&number, "ptr"},
	} {
		if got := kindOf(tc.value); got != tc.want {
			t.Errorf("kindOf a %T = %q, want %q", tc.value, got, tc.want)
		}
	}
}
