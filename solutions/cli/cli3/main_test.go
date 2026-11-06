// cli3
// Make the tests pass!

package main_test

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"testing"
)

// run is the whole program, without os.Args, os.Stdout and os.Exit, so a test
// can call it. It returns the exit code: 0 for success, 2 for a wrong command line.
func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("tool", flag.ContinueOnError)
	fs.SetOutput(stderr)
	count := fs.Int("n", 1, "how many times to say hello")
	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 2
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
