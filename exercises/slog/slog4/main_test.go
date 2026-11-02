// slog4
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"bytes"
	"io"
	"log/slog"
	"testing"
)

// newLogger writes text records to w, without the time, and with the value of
// every "password" attribute replaced by [hidden].
func newLogger(w io.Writer) *slog.Logger {
	// TODO: Give the handler a ReplaceAttr function: drop the attribute with the key slog.TimeKey, and hide the password.
	return slog.New(slog.NewTextHandler(w, nil))
}

func TestLoggerHidesPasswordsAndTime(t *testing.T) {
	var out bytes.Buffer

	newLogger(&out).Info("login", "user", "ana", "password", "s3cret")

	want := "level=INFO msg=login user=ana password=[hidden]\n"
	if out.String() != want {
		t.Errorf("logged %q, want %q", out.String(), want)
	}
}
