// encoding1
// Make the tests pass!

package main_test

import (
	"encoding/base64"
	"testing"
)

// token writes the bytes in a form that is safe in a URL: no +, / or = in it.
func token(data []byte) string {
	return base64.RawURLEncoding.EncodeToString(data)
}

func TestToken(t *testing.T) {
	got := token([]byte{0xfb, 0xff})

	if want := "-_8"; got != want {
		t.Errorf("token = %q, want %q", got, want)
	}
}
