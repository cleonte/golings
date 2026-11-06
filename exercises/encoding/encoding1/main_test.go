// encoding1
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"encoding/base64"
	"testing"
)

// token writes the bytes in a form that is safe in a URL: no +, / or = in it.
func token(data []byte) string {
	// TODO: The standard encoding uses + and / and ends with = padding, none of which belongs in a URL. The package has a URL-safe alphabet, and a form without padding.
	return base64.StdEncoding.EncodeToString(data)
}

func TestToken(t *testing.T) {
	got := token([]byte{0xfb, 0xff})

	if want := "-_8"; got != want {
		t.Errorf("token = %q, want %q", got, want)
	}
}
