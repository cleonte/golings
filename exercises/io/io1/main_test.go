// io1
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"io"
	"strings"
	"testing"
	"testing/iotest"
)

// readAll returns everything r has to give.
func readAll(r io.Reader) (string, error) {
	buf := make([]byte, 8)

	// TODO: A single Read may return fewer bytes than there are left, and the reader isn't finished. Keep reading until the end, or use the function of the io package that does it.
	n, err := r.Read(buf)
	return string(buf[:n]), err
}

func TestReadAllOfASmallReader(t *testing.T) {
	got, err := readAll(strings.NewReader("hello"))
	if err != nil || got != "hello" {
		t.Errorf("readAll = %q, %v, want %q", got, err, "hello")
	}
}

func TestReadAllOfAReaderThatGivesOneByteAtATime(t *testing.T) {
	got, err := readAll(iotest.OneByteReader(strings.NewReader("hello world")))
	if err != nil {
		t.Fatal(err)
	}
	if got != "hello world" {
		t.Errorf("readAll = %q, want %q: a Read can return less than what is left", got, "hello world")
	}
}
