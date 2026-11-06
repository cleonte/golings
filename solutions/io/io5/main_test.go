// io5
// Make the tests pass!

package main_test

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"strings"
	"testing"
)

// copyAndHash copies src to dst and returns the SHA-256 of what it copied, as hex. src can be read only once.
func copyAndHash(dst io.Writer, src io.Reader) (string, error) {
	h := sha256.New()

	if _, err := io.Copy(io.MultiWriter(dst, h), src); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func TestCopyAndHash(t *testing.T) {
	var dst bytes.Buffer

	sum, err := copyAndHash(&dst, strings.NewReader("hello"))

	if err != nil {
		t.Fatal(err)
	}
	if dst.String() != "hello" {
		t.Errorf("copied %q, want %q", dst.String(), "hello")
	}
	if want := "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"; sum != want {
		t.Errorf("hash = %s, want %s", sum, want)
	}
}
