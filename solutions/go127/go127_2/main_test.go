// go127_2
// Make the tests pass!

package main_test

import (
	"testing"
	"uuid"
)

// newID returns a new random identifier, as text.
func newID() string {
	return uuid.New().String()
}

// parseID reads the text form of an identifier. It reports an error when the
// text is not a UUID.
func parseID(text string) (uuid.UUID, error) {
	return uuid.Parse(text)
}

func TestNewID(t *testing.T) {
	first, second := newID(), newID()

	if first == second {
		t.Errorf("two new IDs are both %q, want them to differ", first)
	}
	if _, err := uuid.Parse(first); err != nil {
		t.Errorf("newID() = %q, which is not a UUID: %v", first, err)
	}
}

func TestParseID(t *testing.T) {
	const text = "123e4567-e89b-12d3-a456-426614174000"

	id, err := parseID(text)
	if err != nil {
		t.Fatalf("parseID(%q) failed: %v", text, err)
	}
	if id.String() != text {
		t.Errorf("parseID(%q) gave %s, want the same UUID", text, id)
	}

	if _, err := parseID("not-a-uuid"); err == nil {
		t.Error("parseID(\"not-a-uuid\") worked, want an error")
	}
}
