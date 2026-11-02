// slog3
// Make the tests pass!

// I AM NOT DONE
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

// forOrder returns a logger that adds the order_id to everything it logs.
func forOrder(logger *slog.Logger, id int) *slog.Logger {
	// TODO: Return a logger with the pair "order_id" and the id already added.
	return logger
}

// logCustomer records the customer's name and whether they are a VIP, grouped
// under "customer".
func logCustomer(logger *slog.Logger, name string, vip bool) {
	// TODO: Put "name" and "vip" in a slog.Group called "customer".
	logger.Info("customer", "name", name, "vip", vip)
}

func TestForOrder(t *testing.T) {
	var out bytes.Buffer
	logger := forOrder(newLogger(&out), 7)

	logger.Info("created")
	logger.Info("paid")

	want := "level=INFO msg=created order_id=7\nlevel=INFO msg=paid order_id=7\n"
	if out.String() != want {
		t.Errorf("logged %q, want %q", out.String(), want)
	}
}

func TestLogCustomer(t *testing.T) {
	var out bytes.Buffer

	logCustomer(newLogger(&out), "ana", true)

	want := "level=INFO msg=customer customer.name=ana customer.vip=true\n"
	if out.String() != want {
		t.Errorf("logged %q, want %q", out.String(), want)
	}
}
