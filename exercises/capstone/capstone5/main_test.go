// capstone5
// Make the tests pass!

// I AM NOT DONE
package main

import "testing"

func TestMergeAddsTheCounts(t *testing.T) {
	a, b := NewCounter(), NewCounter()
	a.Add("go go rust")
	b.Add("go zig")

	a.Merge(b)
	if a.Count("go") != 3 || a.Count("rust") != 1 || a.Count("zig") != 1 {
		t.Error("Merge did not add the counts of a word that is in both counters")
	}
}

func TestCountAllAddsUpEveryText(t *testing.T) {
	c := CountAll([]string{"go go", "go rust", "rust", "zig go"})

	if c.Count("go") != 4 || c.Count("rust") != 2 || c.Total() != 7 {
		t.Error("CountAll lost counts: the counts of every text have to add up")
	}
}
