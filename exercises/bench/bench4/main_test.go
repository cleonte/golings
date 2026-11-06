// bench4
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"bytes"
	"runtime/pprof"
	"strings"
	"testing"
)

// heapProfile returns the heap profile as text: what is allocated now, and where.
func heapProfile() (string, error) {
	var buf bytes.Buffer
	// TODO: This writes the profile of the goroutines. The profile of the memory has another name: look it up by that name.
	if err := pprof.Lookup("goroutine").WriteTo(&buf, 1); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func TestHeapProfileIsAHeapProfile(t *testing.T) {
	text, err := heapProfile()

	if err != nil {
		t.Fatalf("heapProfile returned an error: %v", err)
	}
	if !strings.HasPrefix(text, "heap profile:") {
		first, _, _ := strings.Cut(text, "\n")
		t.Errorf("the profile starts with %q, want a line that starts with %q", first, "heap profile:")
	}
}
