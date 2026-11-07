package main

import (
	"flag"
	"io"

	"github.com/cleonte/golings/exercises/capstone/capstone4/store"
)

// run is the whole program: the arguments and the streams come in as
// parameters, and the exit code comes out.
func run(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("wc", flag.ContinueOnError)
	fs.SetOutput(stderr)
	total := fs.Bool("total", false, "print only the number of words")

	// TODO: Parse args (a wrong flag is exit code 2), count the words of stdin with AddFrom (a read error is printed to stderr and exit code 1), then print the total if -total is set, or else save the counts with store.NewText(stdout).
	_, _ = total, store.NewText
	return 0
}
