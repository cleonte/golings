// iterators2
// Make me compile!

// I AM NOT DONE
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
	// TODO: A Seq gives one value per round, so there is only one variable to write.
	for position, fruit := range fruits() {
		fmt.Println(position, fruit)
	}

	// A Seq2 gives two values: here the position and the fruit.
	for position, fruit := range slices.All([]string{"apple", "banana", "cherry"}) {
		fmt.Println(position, fruit)
	}
}
