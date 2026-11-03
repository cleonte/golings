// pitfalls4
// Make the tests pass!

package main_test

import "testing"

// countWords says how many times each word appears.
func countWords(words []string) map[string]int {
	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}
	return counts
}

func TestCountWords(t *testing.T) {
	got := countWords([]string{"go", "is", "fun", "go"})

	if got["go"] != 2 || got["fun"] != 1 || len(got) != 3 {
		t.Errorf("countWords gave %v, want go:2 is:1 fun:1", got)
	}
}

func TestReadingANilMapIsFine(t *testing.T) {
	var counts map[string]int

	if counts["go"] != 0 || len(counts) != 0 {
		t.Error("a nil map should read as empty")
	}
}
