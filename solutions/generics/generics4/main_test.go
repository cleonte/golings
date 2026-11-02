// generics4
// Make the tests pass!

package main_test

import "testing"

// Stack is a last in, first out collection of values of one type.
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes the item on top and returns it. An empty stack gives the zero
// value of T and false.
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	last := len(s.items) - 1
	item := s.items[last]
	s.items = s.items[:last]
	return item, true
}

func TestStackOfInts(t *testing.T) {
	var s Stack[int]
	s.Push(1)
	s.Push(2)
	s.Push(3)

	for _, want := range []int{3, 2, 1} {
		if got, ok := s.Pop(); !ok || got != want {
			t.Errorf("Pop() = %d, %t, want %d, true", got, ok, want)
		}
	}
}

func TestStackOfStrings(t *testing.T) {
	var s Stack[string]
	s.Push("a")
	s.Push("b")

	if got, ok := s.Pop(); !ok || got != "b" {
		t.Errorf("Pop() = %q, %t, want \"b\", true", got, ok)
	}
}

func TestEmptyStackGivesTheZeroValue(t *testing.T) {
	var s Stack[string]

	if got, ok := s.Pop(); ok || got != "" {
		t.Errorf("Pop() = %q, %t, want \"\", false", got, ok)
	}
}
