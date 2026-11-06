// io1
// Make the tests pass!

package main_test

import (
	"io"
	"strings"
	"testing"
	"testing/iotest"
)

// readAll returns everything r has to give.
func readAll(r io.Reader) (string, error) {
	data, err := io.ReadAll(r)
	return string(data), err
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
