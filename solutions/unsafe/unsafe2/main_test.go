// unsafe2
// Make the tests pass!

package main_test

import (
	"testing"
	"unsafe"
)

// view returns the text of b without copying it: the string and the slice
// share the same bytes.
func view(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

func TestViewSharesTheBytes(t *testing.T) {
	b := []byte("hello")
	s := view(b)

	if s != "hello" {
		t.Fatalf("view(b) = %q, want %q", s, "hello")
	}
	// A string must never change, so this is only safe in a test that proves
	// the memory is shared: nobody may write to b while s is in use.
	b[0] = 'j'
	if s != "jello" {
		t.Errorf("after changing b, s = %q: view copied the bytes instead of sharing them", s)
	}
}

func TestViewOfNothing(t *testing.T) {
	if s := view(nil); s != "" {
		t.Errorf("view(nil) = %q, want an empty string", s)
	}
}
