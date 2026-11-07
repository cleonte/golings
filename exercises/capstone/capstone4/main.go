// The word counter: main hands the real arguments and streams to run, which a
// test can call with others.
package main

import "os"

func main() {
	os.Exit(run(os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}
