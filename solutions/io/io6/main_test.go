// io6
// Make the tests pass!

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
	n, err := u.r.Read(p)
	for i := range n {
		if c := p[i]; 'a' <= c && c <= 'z' {
			p[i] = c - 'a' + 'A'
		}
	}
	return n, err
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
