// rand1
// Make me compile!

// I AM NOT DONE
package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	// TODO: math/rand/v2 names its functions differently from math/rand: IntN, not Intn. Use the name the compiler suggests.
	roll := rand.Intn(6) + 1

	if roll < 1 || roll > 6 {
		fmt.Println("this die has a face that does not exist:", roll)
		return
	}
	fmt.Println("you rolled a number from 1 to 6")
}
