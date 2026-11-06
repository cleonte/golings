// io3
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"bufio"
	"errors"
	"io"
	"strings"
	"testing"
)

// readLines returns the lines of r, or the error that stopped the reading.
func readLines(r io.Reader) ([]string, error) {
	var lines []string

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// TODO: Scan returns false both at the end of the input and when something went wrong, such as a line that is too long. Ask the scanner which one it was, and return that error.
	return lines, nil
}

func TestReadLines(t *testing.T) {
	got, err := readLines(strings.NewReader("one\ntwo\nthree\n"))
	if err != nil || len(got) != 3 {
		t.Errorf("readLines = %q, %v, want 3 lines and no error", got, err)
	}
}

func TestReadLinesTellsWhenALineIsTooLong(t *testing.T) {
	input := "short\n" + strings.Repeat("x", 100_000) + "\nnever reached\n"

	_, err := readLines(strings.NewReader(input))

	if !errors.Is(err, bufio.ErrTooLong) {
		t.Errorf("readLines returned %v, want bufio.ErrTooLong: a scanner stops at a line over 64 KiB", err)
	}
}
