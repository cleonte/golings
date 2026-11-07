package main

import (
	"errors"
	"flag"
	"fmt"
	"io"

	"github.com/cleonte/golings/exercises/capstone/capstone4/store"
)

// run is the whole program: the arguments and the streams come in as
// parameters, and the exit code comes out.
func run(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("wc", flag.ContinueOnError)
	fs.SetOutput(stderr)
	total := fs.Bool("total", false, "print only the number of words")
	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 2
	}

	c := NewCounter()
	if err := c.AddFrom(stdin); err != nil {
		fmt.Fprintln(stderr, "wc:", err)
		return 1
	}
	if *total {
		fmt.Fprintln(stdout, c.Total())
		return 0
	}
	if err := c.Save(store.NewText(stdout)); err != nil {
		fmt.Fprintln(stderr, "wc:", err)
		return 1
	}
	return 0
}
