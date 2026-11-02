// generics3
// Make me compile!

// I AM NOT DONE
package main

import (
	"fmt"
)

type Celsius float64

type Number interface {
	// TODO: Celsius is its own type, built on float64. Accept every type whose underlying type is one of these.
	int | float64
}

func Sum[T Number](values []T) T {
	var total T
	for _, value := range values {
		total += value
	}
	return total
}

// TODO: Comparing with > needs a constraint that allows it. The cmp package has one for the types that can be ordered.
func Biggest[T any](values []T) T {
	biggest := values[0]
	for _, value := range values[1:] {
		if value > biggest {
			biggest = value
		}
	}
	return biggest
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3}))
	fmt.Println(Sum([]Celsius{21.5, 19, 23.25}))
	fmt.Println(Biggest([]string{"pear", "apple", "fig"}))
	fmt.Println(Biggest([]float64{1.5, 0.5}))
}
