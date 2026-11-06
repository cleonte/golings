// encoding2
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"encoding/hex"
	"testing"
)

// fromHex decodes a hexadecimal text, or says why it can't.
func fromHex(text string) ([]byte, error) {
	// TODO: The error is thrown away, so a text that isn't hexadecimal looks like a success. Return the error with the bytes.
	data, _ := hex.DecodeString(text)
	return data, nil
}

func TestFromHex(t *testing.T) {
	got, err := fromHex("cafe")

	if err != nil || len(got) != 2 || got[0] != 0xca || got[1] != 0xfe {
		t.Errorf("fromHex(\"cafe\") = %x, %v, want cafe and no error", got, err)
	}
}

func TestFromHexRefusesTextThatIsNotHex(t *testing.T) {
	for _, text := range []string{"zz", "abc"} {
		if _, err := fromHex(text); err == nil {
			t.Errorf("fromHex(%q) returned no error", text)
		}
	}
}
