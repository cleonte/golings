// capstone4
// Make the tests pass!

// I AM NOT DONE
package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
	"testing/iotest"
)

func TestRunPrintsTheCounts(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := run(nil, strings.NewReader("b a b"), &stdout, &stderr)
	if code != 0 || stdout.String() != "a 1\nb 2\n" {
		t.Errorf("exit code %d and stdout %q, want 0 and %q", code, stdout.String(), "a 1\nb 2\n")
	}
}

func TestRunTotalPrintsOnlyTheTotal(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := run([]string{"-total"}, strings.NewReader("b a b"), &stdout, &stderr)
	if code != 0 || stdout.String() != "3\n" {
		t.Errorf("exit code %d and stdout %q, want 0 and %q", code, stdout.String(), "3\n")
	}
}

func TestRunWrongFlagIsExitCodeTwo(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := run([]string{"-nope"}, strings.NewReader("b a b"), &stdout, &stderr)
	if code != 2 {
		t.Errorf("exit code %d, want 2", code)
	}
	if stdout.Len() != 0 || stderr.Len() == 0 {
		t.Errorf("a wrong flag is reported on stderr, and nothing goes to stdout (stdout %q)", stdout.String())
	}
}

func TestRunReadErrorIsExitCodeOne(t *testing.T) {
	var stdout, stderr bytes.Buffer

	code := run(nil, iotest.ErrReader(errors.New("disk gone")), &stdout, &stderr)
	if code != 1 || !strings.Contains(stderr.String(), "disk gone") {
		t.Errorf("exit code %d and stderr %q, want 1 and the error of the reader", code, stderr.String())
	}
}
