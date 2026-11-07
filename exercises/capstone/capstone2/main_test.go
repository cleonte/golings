// capstone2
// Make the tests pass!

// I AM NOT DONE
package main

import (
	"errors"
	"strings"
	"testing"
	"testing/iotest"
)

func TestAddFromCountsALongLine(t *testing.T) {
	c := NewCounter()
	// One line of 120000 bytes, longer than a bufio.Scanner takes as a line.
	if err := c.AddFrom(strings.NewReader(strings.Repeat("go ", 40000))); err != nil {
		t.Fatalf("AddFrom failed: %v", err)
	}
	if got := c.Total(); got != 40000 {
		t.Errorf("Total = %d, want 40000", got)
	}
}

func TestAddFromCountsWordsAcrossLines(t *testing.T) {
	c := NewCounter()
	if err := c.AddFrom(strings.NewReader("one two\nthree Two\n")); err != nil {
		t.Fatalf("AddFrom failed: %v", err)
	}
	if got := c.Count("two"); got != 2 {
		t.Errorf("Count(two) = %d, want 2", got)
	}
}

func TestAddFromReportsAReadError(t *testing.T) {
	c := NewCounter()
	if err := c.AddFrom(iotest.ErrReader(errors.New("disk gone"))); err == nil {
		t.Error("AddFrom dropped the error of the reader")
	}
}
