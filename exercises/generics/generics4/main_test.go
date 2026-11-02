// generics4
// Make the tests pass!

// I AM NOT DONE
package main_test

import "testing"

// Stack is a last in, first out collection of values of one type.
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	// TODO: Put the item on top of the stack.
}

// Pop removes the item on top and returns it. An empty stack gives the zero
// value of T and false.
func (s *Stack[T]) Pop() (T, bool) {
	// TODO: Take the last item off the stack. For an empty stack, return the zero value of T.
	var zero T
	return zero, false
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
