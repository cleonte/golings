// slog4
// Make the tests pass!

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
	return slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case slog.TimeKey:
				return slog.Attr{}
			case "password":
				return slog.String("password", "[hidden]")
			}
			return a
		},
	}))
}

func TestLoggerHidesPasswordsAndTime(t *testing.T) {
	var out bytes.Buffer

	newLogger(&out).Info("login", "user", "ana", "password", "s3cret")

	want := "level=INFO msg=login user=ana password=[hidden]\n"
	if out.String() != want {
		t.Errorf("logged %q, want %q", out.String(), want)
	}
}
