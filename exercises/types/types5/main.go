// types5
// Make me compile!

// I AM NOT DONE
package main

import (
	"fmt"
)

type Meters float64
type Feet float64

func main() {
	height := Meters(2)
	extra := Feet(3)

	// TODO: Meters and Feet are different types, even though both are float64 underneath. Convert extra to Meters (a foot is 0.3048 metres) before adding it to height.
	total := height + extra

	fmt.Printf("total: %.2f m\n", total)
}
