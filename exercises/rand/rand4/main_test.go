// rand4
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"strings"
	"testing"
)

// newToken returns a secret token for a login.
func newToken() string {
	// TODO: Anything made by math/rand can be predicted by someone who knows a few values, so it is not for secrets. The crypto/rand package has a function that returns random text (Go 1.24).
	return strings.Repeat("x", 26)
}

func TestTokenIs26CharactersOfBase32(t *testing.T) {
	token := newToken()

	if len(token) != 26 {
		t.Fatalf("the token %q has %d characters, want 26", token, len(token))
	}
	if strings.Trim(token, "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567") != "" {
		t.Errorf("the token %q has characters outside the base32 alphabet", token)
	}
}

func TestTwoTokensDiffer(t *testing.T) {
	if newToken() == newToken() {
		t.Error("two tokens are the same: they have to be random")
	}
}
