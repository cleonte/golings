// errors6
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"errors"
	"fmt"
	"testing"
)

type RateLimitError struct {
	WaitSeconds int
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("rate limited, try again in %d seconds", e.WaitSeconds)
}

// waitTime says how long to wait before trying again, when err (or an error it
// wraps) is a *RateLimitError.
func waitTime(err error) (int, bool) {
	// TODO: Look for a *RateLimitError in the chain of err with errors.AsType (Go 1.26), and return its WaitSeconds.
	return 0, false
}

func TestWaitTime(t *testing.T) {
	limited := &RateLimitError{WaitSeconds: 30}

	tests := []struct {
		name     string
		err      error
		wantWait int
		wantOK   bool
	}{
		{name: "direct", err: limited, wantWait: 30, wantOK: true},
		{name: "wrapped", err: fmt.Errorf("calling the API: %w", limited), wantWait: 30, wantOK: true},
		{name: "joined", err: errors.Join(errors.New("first"), limited), wantWait: 30, wantOK: true},
		{name: "another error", err: errors.New("boom")},
		{name: "no error"},
	}

	for _, tc := range tests {
		wait, ok := waitTime(tc.err)
		if wait != tc.wantWait || ok != tc.wantOK {
			t.Errorf("%s: waitTime() = %d, %t, want %d, %t", tc.name, wait, ok, tc.wantWait, tc.wantOK)
		}
	}
}
