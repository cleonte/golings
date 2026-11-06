// layout5
// Make me compile!

// I AM NOT DONE
package main

import (
	"fmt"

	"github.com/cleonte/golings/exercises/layout/layout5/store/v2"
)

func main() {
	// TODO: The last part of an import path is the name of a folder, and the name of the package is what the package clause of its files says: this one is called store, not v2. Use that name.
	s := v2.New()
	s.Set("language", "Go")
	fmt.Println(s.Get("language"))
}
