// generics6
// Make me compile!

package main

import "fmt"

type Set[T comparable] = map[T]struct{}

func add[T comparable](set Set[T], value T) {
	set[value] = struct{}{}
}

func main() {
	names := Set[string]{}
	add(names, "ana")
	add(names, "carl")
	add(names, "ana")

	// Set[string] is the same type as map[string]struct{}: no conversion needed.
	var plain map[string]struct{} = names
	fmt.Println(len(plain), "different names")
}
