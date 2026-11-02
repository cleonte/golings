// concurrent11
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"slices"
	"testing"
	"time"
)

// squares computes the squares in parallel: one goroutine for each number.
func squares(numbers []int) []int {
	results := make([]int, len(numbers))

	for i, n := range numbers {
		// TODO: Start each goroutine with a sync.WaitGroup's Go method (Go 1.25), and wait for all of them before returning.
		go func() {
			time.Sleep(10 * time.Millisecond) // a slow calculation
			results[i] = n * n
		}()
	}

	return results
}

func TestSquares(t *testing.T) {
	got := squares([]int{1, 2, 3, 4, 5})

	if want := []int{1, 4, 9, 16, 25}; !slices.Equal(got, want) {
		t.Errorf("squares gave %v, want %v", got, want)
	}
}
