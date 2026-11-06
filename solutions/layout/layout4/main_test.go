// layout4
// Make the tests pass!

package layout4_test

import (
	"testing"

	"github.com/cleonte/golings/exercises/layout/layout4/registry"

	_ "github.com/cleonte/golings/exercises/layout/layout4/plugins/english"
	_ "github.com/cleonte/golings/exercises/layout/layout4/plugins/french"
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
