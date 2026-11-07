// capstone1
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

// The first step of the word counter: the Counter type is in counter.go, and
// main uses it. One method that main calls is missing.
func main() {
	c := NewCounter()
	c.Add("the quick brown fox jumps over the lazy dog")
	c.Add("The end")

	fmt.Println("words:", c.Total())
	fmt.Println("the:", c.Count("the"))
}
