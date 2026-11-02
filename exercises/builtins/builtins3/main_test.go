// builtins3
// Make the tests pass!

// I AM NOT DONE
package main_test

import "testing"

// wipe overwrites the secret, so it doesn't stay around in memory. The slice
// keeps its length.
func wipe(secret []byte) {
	// TODO: Set every byte to zero. The built-in clear does it for a slice too.
}

func TestWipe(t *testing.T) {
	secret := []byte("hunter2")
	start := secret[:3]

	wipe(secret)

	if len(secret) != 7 {
		t.Errorf("len(secret) = %d, want 7: wiping doesn't shorten the slice", len(secret))
	}
	for i, b := range secret {
		if b != 0 {
			t.Errorf("secret[%d] = %d, want 0", i, b)
		}
	}
	if start[0] != 0 {
		t.Error("the bytes in the array have to be overwritten, not just hidden")
	}
}
