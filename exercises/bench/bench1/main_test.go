// bench1
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"testing"
)

// join glues the words together with commas.
func join(words []string) string {
	// TODO: += builds a whole new string at every step, so every word allocates. Use a strings.Builder (call Grow first, since you know about how big the result is) or strings.Join.
	result := ""
	for i, word := range words {
		if i > 0 {
			result += ","
		}
		result += word
	}
	return result
}

var sink string

func TestJoinAllocatesOnce(t *testing.T) {
	words := make([]string, 50)
	for i := range words {
		words[i] = "word"
	}
	if got := join(words); len(got) != 50*4+49 {
		t.Fatalf("join gave %d characters, want %d", len(got), 50*4+49)
	}

	// AllocsPerRun counts the memory allocations of a function: it needs no clock.
	allocs := testing.AllocsPerRun(10, func() { sink = join(words) })

	if allocs > 2 {
		t.Errorf("join allocates more than twice per call, want 2 or fewer")
	}
}
