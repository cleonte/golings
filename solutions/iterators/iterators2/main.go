// iterators2
// Make me compile!

package main

import (
	"fmt"
	"iter"
	"slices"
)

func fruits() iter.Seq[string] {
	return slices.Values([]string{"apple", "banana", "cherry"})
}

func main() {
	for fruit := range fruits() {
		fmt.Println(fruit)
	}

	// A Seq2 gives two values: here the position and the fruit.
	for position, fruit := range slices.All([]string{"apple", "banana", "cherry"}) {
		fmt.Println(position, fruit)
	}
}
