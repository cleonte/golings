// bench5
// Make the tests pass!

package main_test

import (
	"testing"
)

var calls int

// work is what the benchmark is meant to measure.
func work() {
	calls++
}

func BenchmarkWork(b *testing.B) {
	for b.Loop() {
		work()
	}
}

func TestTheBenchmarkRunsTheWorkInItsLoop(t *testing.T) {
	calls = 0
	// testing.Benchmark runs a benchmark from a test and says how many rounds
	// it ran. It takes about a second.
	result := testing.Benchmark(BenchmarkWork)

	if calls < result.N {
		t.Error("work ran fewer times than the benchmark has rounds: it is not in the loop")
	}
}
