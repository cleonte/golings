// layout5
// Make me compile!

package main

import (
	"fmt"

	"github.com/cleonte/golings/exercises/layout/layout5/store/v2"
)

func main() {
	s := store.New()
	s.Set("language", "Go")
	fmt.Println(s.Get("language"))
}
