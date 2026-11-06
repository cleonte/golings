// cli2
// Make the tests pass!

package main_test

import (
	"flag"
	"io"
	"slices"
	"testing"
)

// parseArgs reads the -v flag and returns the file names that come after the flags.
func parseArgs(args []string) (verbose bool, names []string, err error) {
	fs := flag.NewFlagSet("cat", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.BoolVar(&verbose, "v", false, "print more")
	if err := fs.Parse(args); err != nil {
		return false, nil, err
	}

	return verbose, fs.Args(), nil
}

func TestTheFlagIsNotAFileName(t *testing.T) {
	verbose, names, err := parseArgs([]string{"-v", "a.txt", "b.txt"})

	if err != nil {
		t.Fatalf("parseArgs returned an error: %v", err)
	}
	if !verbose {
		t.Error("-v was not read")
	}
	if want := []string{"a.txt", "b.txt"}; !slices.Equal(names, want) {
		t.Errorf("names = %q, want %q", names, want)
	}
}

func TestFlagsStopAtTheFirstFileName(t *testing.T) {
	// The flag package stops at the first argument that is not a flag, so
	// -v after a file name is just another file name.
	verbose, names, err := parseArgs([]string{"a.txt", "-v"})

	if err != nil {
		t.Fatalf("parseArgs returned an error: %v", err)
	}
	if verbose {
		t.Error("-v after a file name was read as a flag")
	}
	if want := []string{"a.txt", "-v"}; !slices.Equal(names, want) {
		t.Errorf("names = %q, want %q", names, want)
	}
}
