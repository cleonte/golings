// builtins1
// Make me compile!

package main

import "fmt"

func main() {
	morning, noon, evening := 14, 21, 17
	fmt.Println("temperatures:", morning, noon, evening)

	fmt.Println("warmest:", max(morning, noon, evening))
	fmt.Println("coldest:", min(morning, noon, evening))

	// Strings are ordered too: this is the first name in alphabetical order.
	fmt.Println("first:", min("pear", "apple", "fig"))
}
