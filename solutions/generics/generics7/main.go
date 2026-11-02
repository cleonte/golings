// generics7
// Make me compile!

package main

import "fmt"

type Adder[A Adder[A]] interface {
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
