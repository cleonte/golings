// pointers6
// Make me compile!

package main

import "fmt"

// Options uses a pointer for each setting, so that nil can mean "not set".
type Options struct {
	Retries *int
	Verbose *bool
	Label   *string
}

func main() {
	opts := Options{
		Retries: new(3),
		Verbose: new(true),
		Label:   new("fast"),
	}

	fmt.Println("retries:", *opts.Retries)
	fmt.Println("verbose:", *opts.Verbose)
	fmt.Println("label:", *opts.Label)
}
