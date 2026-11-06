// net4
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"testing"
	"time"
)

// isTimeout reports whether err comes from a deadline that has passed, even when
// it is wrapped in other errors.
func isTimeout(err error) bool {
	// TODO: A network read does not return os.ErrDeadlineExceeded itself: it returns a *net.OpError that wraps it, and fmt.Errorf with %w can wrap it again. Look inside the chain with errors.Is.
	return err == os.ErrDeadlineExceeded
}

var _ = errors.Is

func expiredRead(t *testing.T) error {
	t.Helper()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { ln.Close() })
	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { conn.Close() })

	// A deadline in the past makes the next read fail at once with a timeout.
	conn.SetReadDeadline(time.Now().Add(-time.Second))
	_, err = conn.Read(make([]byte, 1))
	return err
}

func TestAReadPastItsDeadlineIsATimeout(t *testing.T) {
	err := expiredRead(t)

	if !isTimeout(err) {
		t.Errorf("isTimeout of a %T from a read past its deadline = false, want true", err)
	}
}

func TestAWrappedTimeoutIsStillATimeout(t *testing.T) {
	err := fmt.Errorf("fetching the page: %w", expiredRead(t))

	if !isTimeout(err) {
		t.Error("isTimeout of a wrapped timeout = false, want true")
	}
}

func TestOtherErrorsAreNotTimeouts(t *testing.T) {
	if isTimeout(io.EOF) {
		t.Error("isTimeout(io.EOF) = true, want false")
	}
}
