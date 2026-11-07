// capstone3
// Make the tests pass!

// I AM NOT DONE
package main

import (
	"bytes"
	"errors"
	"testing"

	"github.com/cleonte/golings/exercises/capstone/capstone3/store"
)

// fakeSaver remembers what it was asked to save, and can fail on demand.
type fakeSaver struct {
	saved []map[string]int
	err   error
}

func (f *fakeSaver) Save(counts map[string]int) error {
	f.saved = append(f.saved, counts)
	return f.err
}

func TestSaveHandsTheCountsToTheSaver(t *testing.T) {
	c := NewCounter()
	c.Add("b a b")
	fake := &fakeSaver{}

	if err := c.Save(fake); err != nil {
		t.Fatalf("Save failed: %v", err)
	}
	if len(fake.saved) != 1 || fake.saved[0]["b"] != 2 {
		t.Errorf("the saver was not given the counts: %v", fake.saved)
	}
}

func TestSaveReturnsTheErrorOfTheSaver(t *testing.T) {
	c := NewCounter()
	fake := &fakeSaver{err: errors.New("disk full")}

	if err := c.Save(fake); !errors.Is(err, fake.err) {
		t.Errorf("Save returned %v, want the error of the saver", err)
	}
}

func TestSaveWithTheTextStore(t *testing.T) {
	c := NewCounter()
	c.Add("b a b")
	var out bytes.Buffer

	if err := c.Save(store.NewText(&out)); err != nil {
		t.Fatalf("Save failed: %v", err)
	}
	if want := "a 1\nb 2\n"; out.String() != want {
		t.Errorf("the store wrote %q, want %q", out.String(), want)
	}
}
