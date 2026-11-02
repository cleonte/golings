// generics5
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

// Index returns the position of target in items, or -1 when it is not there.
// TODO: Comparing with == needs a constraint that says the values can be compared.
func Index[T any](items []T, target T) int {
	for i, item := range items {
		if item == target {
			return i
		}
	}
	return -1
}

// Count says how many times each value appears.
// TODO: The same goes for map keys: T has to be comparable.
func Count[T any](items []T) map[T]int {
	counts := make(map[T]int)
	for _, item := range items {
		counts[item]++
	}
	return counts
}

func main() {
	names := []string{"ana", "carl", "ana", "maria"}

	fmt.Println(Index(names, "carl"))
	fmt.Println(Count(names))
	fmt.Println(Index([]float64{0.5, 1.5}, 1.5))
}
