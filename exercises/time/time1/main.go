// time1
// Make me compile!

// I AM NOT DONE
package main

import (
	"fmt"
	"time"
)

func main() {
	retries := 3

	// TODO: A time.Duration is not a plain number, so Go won't multiply an int by time.Second. Convert the number to a time.Duration first.
	wait := retries * time.Second

	fmt.Println("waiting", wait, "before giving up")
	fmt.Println("that is", wait.Seconds(), "seconds")
}
