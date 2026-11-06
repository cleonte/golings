// packages8
// Make the tests pass!

// I AM NOT DONE
package greetings

import "testing"

func TestEnglish(t *testing.T) {
	if got := Greet("en"); got != "Hello" {
		t.Errorf("Greet(en) = %q, want %q", got, "Hello")
	}
}

// TODO: Add a new file to this folder, greeting_fr.go, in the same package. Like greeting_en.go it registers its language in an init function: "fr" greets with "Bonjour".
func TestFrench(t *testing.T) {
	if got := Greet("fr"); got != "Bonjour" {
		t.Errorf("Greet(fr) = %q, want %q", got, "Bonjour")
	}
}
