// cli3
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"testing"
)

// run is the whole program, without os.Args, os.Stdout and os.Exit, so a test
// can call it. It returns the exit code: 0 for success, 2 for a wrong command line.
func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("tool", flag.ContinueOnError)
	// TODO: Messages about the usage go to the error output, not to stdout, and a wrong flag is exit code 2 (code 1 is for a program that failed). Also, -h is not an error: it prints the usage and exits with 0, and Parse tells it apart by returning flag.ErrHelp.
	fs.SetOutput(stdout)
	count := fs.Int("n", 1, "how many times to say hello")
	if err := fs.Parse(args); err != nil {
		return 1
	}

	for range *count {
		fmt.Fprintln(stdout, "hello")
	}
	return 0
}

func TestRunSaysHello(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := run([]string{"-n", "2"}, &stdout, &stderr)

	if code != 0 || stdout.String() != "hello\nhello\n" {
		t.Errorf("run = %d with the output %q, want 0 and two lines of hello", code, stdout.String())
	}
}

func TestAWrongFlagIsExitCode2AndGoesToStderr(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := run([]string{"-bogus"}, &stdout, &stderr)

	if code != 2 {
		t.Errorf("exit code = %d, want 2", code)
	}
	if stdout.Len() != 0 {
		t.Error("the message about the usage went to stdout: it belongs on stderr")
	}
	if stderr.Len() == 0 {
		t.Error("nothing was written to stderr")
	}
}

func TestHelpIsNotAFailure(t *testing.T) {
	var stdout, stderr bytes.Buffer

	if code := run([]string{"-h"}, &stdout, &stderr); code != 0 {
		t.Errorf("exit code of -h = %d, want 0", code)
	}
}
