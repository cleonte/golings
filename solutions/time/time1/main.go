// time1
// Make me compile!

package main

import (
	"fmt"
	"time"
)

func main() {
	retries := 3

	wait := time.Duration(retries) * time.Second

	fmt.Println("waiting", wait, "before giving up")
	fmt.Println("that is", wait.Seconds(), "seconds")
}
