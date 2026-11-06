// bench5
// Make the tests pass!

// I AM NOT DONE
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
	// TODO: work runs once, before the loop, so the benchmark measures nothing. Run it for every round of the loop: the loop is for b.Loop() { ... } and the code to measure goes in its body.
	work()
	for b.Loop() {
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
