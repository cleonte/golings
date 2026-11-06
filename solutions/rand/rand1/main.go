// rand1
// Make me compile!

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	roll := rand.IntN(6) + 1

	if roll < 1 || roll > 6 {
		fmt.Println("this die has a face that does not exist:", roll)
		return
	}
	fmt.Println("you rolled a number from 1 to 6")
}
