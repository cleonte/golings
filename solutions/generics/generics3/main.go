// generics3
// Make me compile!

package main

import (
	"cmp"
	"fmt"
)

type Celsius float64

type Number interface {
	~int | ~float64
}

func Sum[T Number](values []T) T {
	var total T
	for _, value := range values {
		total += value
	}
	return total
}

func Biggest[T cmp.Ordered](values []T) T {
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
