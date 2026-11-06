// packages5
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

// The files of one folder are one package, so main can use shout from
// helpers.go without importing anything.
func main() {
	fmt.Println(shout("all the files of a folder are one package"))
}
