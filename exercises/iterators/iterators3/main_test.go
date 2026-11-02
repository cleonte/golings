// iterators3
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"fmt"
	"iter"
	"testing"
)

// numbered gives each name with its number, counting from 1.
func numbered(names []string) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		// TODO: Pass the number (1 for the first name) and the name to yield.
	}
}

func TestNumbered(t *testing.T) {
	var got []string
	for number, name := range numbered([]string{"ana", "carl", "maria"}) {
		got = append(got, fmt.Sprintf("%d:%s", number, name))
	}

	want := "[1:ana 2:carl 3:maria]"
	if fmt.Sprint(got) != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
