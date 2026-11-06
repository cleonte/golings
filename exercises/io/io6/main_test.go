// io6
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"io"
	"strings"
	"testing"
	"testing/iotest"
)

// upperReader reads from r and turns the letters a to z that it reads into capitals.
type upperReader struct {
	r io.Reader
}

func (u upperReader) Read(p []byte) (int, error) {
	// TODO: This reads nothing and says the input is over. Read from u.r into p, change the n bytes it returned (and only those), and return n and the error it gave, unchanged.
	return 0, io.EOF
}

func TestUpperReader(t *testing.T) {
	got, err := io.ReadAll(upperReader{strings.NewReader("hello, reader")})

	if err != nil || string(got) != "HELLO, READER" {
		t.Errorf("got %q, %v, want %q", got, err, "HELLO, READER")
	}
}

func TestUpperReaderWithSmallReads(t *testing.T) {
	r := upperReader{iotest.OneByteReader(strings.NewReader("abc"))}

	got, err := io.ReadAll(r)

	if err != nil || string(got) != "ABC" {
		t.Errorf("got %q, %v, want %q", got, err, "ABC")
	}
}
