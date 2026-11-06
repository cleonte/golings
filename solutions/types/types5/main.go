// types5
// Make me compile!

package main

import (
	"fmt"
)

type Meters float64
type Feet float64

func main() {
	height := Meters(2)
	extra := Feet(3)

	total := height + Meters(extra*0.3048)

	fmt.Printf("total: %.2f m\n", total)
}
