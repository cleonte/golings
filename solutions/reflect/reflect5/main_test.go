// reflect5
// Make the tests pass!

package main_test

import (
	"reflect"
	"testing"
)

// sum adds up every int inside a value, however deep: in struct fields, slices
// and arrays. Other values count for nothing.
func sum(v reflect.Value) int {
	switch v.Kind() {
	case reflect.Int:
		return int(v.Int())
	case reflect.Struct:
		total := 0
		for i := range v.NumField() {
			total += sum(v.Field(i))
		}
		return total
	case reflect.Slice, reflect.Array:
		total := 0
		for i := range v.Len() {
			total += sum(v.Index(i))
		}
		return total
	}
	return 0
}

type Box struct {
	Label string
	Items int
}

type Warehouse struct {
	Name   string
	Stock  []int
	Boxes  []Box
	Shelf  [2]int
	Closed bool
}

func TestSumWalksTheWholeValue(t *testing.T) {
	w := Warehouse{
		Name:   "north",
		Stock:  []int{1, 2, 3},
		Boxes:  []Box{{"a", 10}, {"b", 20}},
		Shelf:  [2]int{100, 200},
		Closed: true,
	}

	if got, want := sum(reflect.ValueOf(w)), 336; got != want {
		t.Errorf("sum(w) = %d, want %d", got, want)
	}
}
