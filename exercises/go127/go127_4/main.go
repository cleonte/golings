//go:build go1.27

// go127_4
// Make me compile!

// I AM NOT DONE
package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
	"time"
)

type List[T any] struct {
	items []T
}

// Map returns a list with f applied to every item. Its result has another type.
// TODO: Map uses a type U that does not exist. Declare it as a type parameter of the method, after its name (Go 1.27).
func (l List[T]) Map(f func(T) U) List[U] {
	var result List[U]
	for _, item := range l.items {
		result.items = append(result.items, f(item))
	}
	return result
}

func main() {
	names := List[string]{items: []string{"ana", "carl", "maria"}}

	fmt.Println(names.Map(strings.ToUpper).items)
	fmt.Println(names.Map(func(name string) int { return len(name) }).items)

	// The standard library has a generic method too: N works for any integer type.
	r := rand.New(rand.NewPCG(1, 2))
	wait := r.N(10 * time.Second)
	fmt.Println("a wait under ten seconds:", wait < 10*time.Second)
}
