// layout4
// Make the tests pass!

// I AM NOT DONE
package layout4_test

import (
	"testing"

	"github.com/cleonte/golings/exercises/layout/layout4/registry"

	_ "github.com/cleonte/golings/exercises/layout/layout4/plugins/english"
	// TODO: A package that is imported only for what its init function does is imported with the blank identifier _. The french plugin registers "fr" in its init, but nothing imports it yet. Add that import here, like the line above.
)

func TestEnglishIsRegistered(t *testing.T) {
	if got := registry.Greet("en"); got != "Hello" {
		t.Errorf("Greet(en) = %q, want %q", got, "Hello")
	}
}

func TestFrenchIsRegistered(t *testing.T) {
	if got := registry.Greet("fr"); got != "Bonjour" {
		t.Errorf("Greet(fr) = %q, want %q", got, "Bonjour")
	}
}
