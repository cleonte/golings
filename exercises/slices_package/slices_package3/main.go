// slices_package3
// Make me compile!

// I AM NOT DONE
package main

import (
	"fmt"
	"slices"
)

func main() {
	guests := []string{"ana", "carl", "maria", "john"}

	// TODO: Ask slices.Contains whether "maria" is a guest. It needs the slice and the value.
	fmt.Println("maria is invited:", slices.Contains(guests))

	// TODO: Ask slices.Index for the position of "john". It needs the slice and the value.
	fmt.Println("john is at position", slices.Index(guests))

	// A function says what a match is: the first name that is longer than four letters.
	fmt.Println("first long name at", slices.IndexFunc(guests, func(name string) bool { return len(name) > 4 }))
}
