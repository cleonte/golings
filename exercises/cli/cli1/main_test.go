// cli1
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"flag"
	"fmt"
	"io"
	"testing"
)

// greeting builds the greeting that the command line asks for: -name Ann.
func greeting(args []string) (string, error) {
	fs := flag.NewFlagSet("greet", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	name := fs.String("name", "world", "who to greet")

	// TODO: The flag is declared but the arguments are never read, so name always has its default. Parse args with the flag set and return its error.
	return fmt.Sprintf("Hello, %s!", *name), nil
}

// The functions get the arguments as a parameter, not from os.Args, so a test
// can call them with any command line.
func TestTheNameFlagIsRead(t *testing.T) {
	got, err := greeting([]string{"-name", "Ann"})

	if err != nil || got != "Hello, Ann!" {
		t.Errorf(`greeting("-name", "Ann") = %q, %v, want "Hello, Ann!"`, got, err)
	}
}

func TestTheDefaultIsUsedWithoutTheFlag(t *testing.T) {
	got, err := greeting(nil)

	if err != nil || got != "Hello, world!" {
		t.Errorf("greeting() = %q, %v, want %q", got, err, "Hello, world!")
	}
}

func TestAnUnknownFlagIsAnError(t *testing.T) {
	if _, err := greeting([]string{"-colour", "red"}); err == nil {
		t.Error("greeting with an unknown flag returned no error")
	}
}
