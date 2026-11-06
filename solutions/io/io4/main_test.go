// io4
// Make the tests pass!

package main_test

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

// copyAtMost copies at most max bytes from src to dst, and says how many it copied.
func copyAtMost(dst io.Writer, src io.Reader, max int64) (int64, error) {
	return io.Copy(dst, io.LimitReader(src, max))
}

func TestCopyAtMost(t *testing.T) {
	var dst bytes.Buffer

	n, err := copyAtMost(&dst, strings.NewReader("hello world"), 5)

	if err != nil {
		t.Fatal(err)
	}
	if n != 5 || dst.String() != "hello" {
		t.Errorf("copied %d bytes, %q; want 5 bytes, %q", n, dst.String(), "hello")
	}
}

func TestCopyAtMostWithLessThanTheLimit(t *testing.T) {
	var dst bytes.Buffer

	n, _ := copyAtMost(&dst, strings.NewReader("hi"), 5)

	if n != 2 || dst.String() != "hi" {
		t.Errorf("copied %d bytes, %q; want 2 bytes, %q", n, dst.String(), "hi")
	}
}
