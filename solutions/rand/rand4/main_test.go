// rand4
// Make the tests pass!

package main_test

import (
	"crypto/rand"
	"strings"
	"testing"
)

// newToken returns a secret token for a login.
func newToken() string {
	return rand.Text()
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
