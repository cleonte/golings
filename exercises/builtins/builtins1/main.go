// builtins1
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

func main() {
	morning, noon, evening := 14, 21, 17
	fmt.Println("temperatures:", morning, noon, evening)

	// TODO: Give max and min the temperatures to compare. They take two or more values.
	fmt.Println("warmest:", max())
	fmt.Println("coldest:", min())

	// Strings are ordered too: this is the first name in alphabetical order.
	fmt.Println("first:", min("pear", "apple", "fig"))
}
