// pointers6
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

// Options uses a pointer for each setting, so that nil can mean "not set".
type Options struct {
	Retries *int
	Verbose *bool
	Label   *string
}

func main() {
	// TODO: You can't take the address of a value like 3: give new the value instead (Go 1.26).
	opts := Options{
		Retries: &3,
		Verbose: &true,
		Label:   &"fast",
	}

	fmt.Println("retries:", *opts.Retries)
	fmt.Println("verbose:", *opts.Verbose)
	fmt.Println("label:", *opts.Label)
}
