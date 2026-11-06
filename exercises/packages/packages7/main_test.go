// packages7
// Make the tests pass!

// I AM NOT DONE
package order

import "testing"

// The init functions of a package run in the order of the file names, which
// is not an order to rely on. Package-level variables are different: they are
// all set up before any init function runs.
func TestTheServerKnowsTheConfig(t *testing.T) {
	if want := "listening on :8080"; server != want {
		t.Errorf("server = %q, want %q", server, want)
	}
}
