// bench4
// Make the tests pass!

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
	if err := pprof.Lookup("heap").WriteTo(&buf, 1); err != nil {
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
