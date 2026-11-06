// cli5
// Make the tests pass!

package main_test

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"testing"
)

// greet is the command "greet -name Ann".
func greet(args []string, stdout io.Writer) int {
	fs := flag.NewFlagSet("greet", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	name := fs.String("name", "world", "who to greet")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	fmt.Fprintf(stdout, "Hello, %s!\n", *name)
	return 0
}

// dispatch picks the command from the first argument, like "git commit".
func dispatch(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "usage: tool <command>")
		return 2
	}
	switch args[0] {
	case "greet":
		return greet(args[1:], stdout)
	default:
		fmt.Fprintf(stderr, "unknown command %q\n", args[0])
		return 2
	}
}

func TestTheCommandGetsItsOwnFlags(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := dispatch([]string{"greet", "-name", "Ann"}, &stdout, &stderr)

	if code != 0 || stdout.String() != "Hello, Ann!\n" {
		t.Errorf("dispatch = %d with the output %q, want 0 and %q", code, stdout.String(), "Hello, Ann!\n")
	}
}

func TestAnUnknownCommandIsExitCode2(t *testing.T) {
	var stdout, stderr bytes.Buffer

	if code := dispatch([]string{"dance"}, &stdout, &stderr); code != 2 {
		t.Errorf("exit code = %d, want 2", code)
	}
	if stderr.Len() == 0 {
		t.Error("nothing was written to stderr")
	}
}
