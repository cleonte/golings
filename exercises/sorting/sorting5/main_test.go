// sorting5
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"container/list"
	"slices"
	"testing"
)

// touch marks the element e of l as used just now: the most recently used element is at the front.
func touch(l *list.List, e *list.Element) {
	// TODO: Nothing moves. The list package has a method that moves an element to the front without taking it out and putting it back.
}

func values(l *list.List) []string {
	var out []string
	for e := l.Front(); e != nil; e = e.Next() {
		out = append(out, e.Value.(string))
	}
	return out
}

func TestTouchMovesToTheFront(t *testing.T) {
	l := list.New()
	l.PushBack("a")
	l.PushBack("b")
	c := l.PushBack("c")

	touch(l, c)

	if got, want := values(l), []string{"c", "a", "b"}; !slices.Equal(got, want) {
		t.Errorf("the list is %q, want %q", got, want)
	}
}
