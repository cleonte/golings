// concurrent11
// Make the tests pass!

package main_test

import (
	"slices"
	"sync"
	"testing"
	"time"
)

// squares computes the squares in parallel: one goroutine for each number.
func squares(numbers []int) []int {
	results := make([]int, len(numbers))
	var wg sync.WaitGroup

	for i, n := range numbers {
		wg.Go(func() {
			time.Sleep(10 * time.Millisecond) // a slow calculation
			results[i] = n * n
		})
	}

	wg.Wait()
	return results
}

func TestSquares(t *testing.T) {
	got := squares([]int{1, 2, 3, 4, 5})

	if want := []int{1, 4, 9, 16, 25}; !slices.Equal(got, want) {
		t.Errorf("squares gave %v, want %v", got, want)
	}
}
