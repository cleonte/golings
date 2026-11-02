// slices_package3
// Make me compile!

package main

import (
	"fmt"
	"slices"
)

func main() {
	guests := []string{"ana", "carl", "maria", "john"}

	fmt.Println("maria is invited:", slices.Contains(guests, "maria"))
	fmt.Println("john is at position", slices.Index(guests, "john"))

	// A function says what a match is: the first name that is longer than four letters.
	fmt.Println("first long name at", slices.IndexFunc(guests, func(name string) bool { return len(name) > 4 }))
}
