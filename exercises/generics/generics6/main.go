// generics6
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

// TODO: Give the alias a type parameter, so that Set[string] means map[string]struct{}. Map keys have to be comparable.
type Set = map[T]struct{}

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
