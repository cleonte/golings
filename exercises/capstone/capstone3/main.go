// The word counter: main counts two lines of text and prints the totals.
package main

import "fmt"

func main() {
	c := NewCounter()
	c.Add("the quick brown fox jumps over the lazy dog")
	c.Add("The end")

	fmt.Println("words:", c.Total())
	fmt.Println("the:", c.Count("the"))
}
