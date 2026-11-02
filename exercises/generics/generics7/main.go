// generics7
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

// TODO: An Adder[A] adds an A to itself, and A has to be an Adder[A] too. Say so in the type parameter list (Go 1.26 allows a type to refer to itself there).
type Adder [A]interface {
	Add(A) A
}

func Sum[A Adder[A]](first A, rest ...A) A {
	total := first
	for _, other := range rest {
		total = total.Add(other)
	}
	return total
}

type Money int

func (m Money) Add(other Money) Money {
	return m + other
}

func main() {
	fmt.Println(Sum(Money(5), 10, 20))
}
