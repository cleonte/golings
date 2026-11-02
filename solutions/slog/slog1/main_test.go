// slog1
// Make the tests pass!

package main_test

import (
	"bytes"
	"io"
	"log/slog"
	"testing"
)

// newLogger writes the records to w as text, without the time so the output is the
// same every run.
func newLogger(w io.Writer) *slog.Logger {
	return slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}
			return a
		},
	}))
}

// logLogin records that a user logged in, and after how many attempts.
func logLogin(logger *slog.Logger, user string, attempts int) {
	logger.Info("user logged in", "user", user, "attempts", attempts)
}

func TestLogLogin(t *testing.T) {
	var out bytes.Buffer

	logLogin(newLogger(&out), "ana", 3)

	want := "level=INFO msg=\"user logged in\" user=ana attempts=3\n"
	if out.String() != want {
		t.Errorf("logged %q, want %q", out.String(), want)
	}
}
