// fakes2
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"errors"
	"testing"
)

type Sender interface {
	Send(to, body string) error
}

// sendWithRetry tries to send up to tries times and returns the last error if
// every try fails.
func sendWithRetry(s Sender, to, body string, tries int) error {
	// TODO: There is only one attempt. Try again after an error, up to tries attempts in all, and stop at the first success.
	return s.Send(to, body)
}

var errBusy = errors.New("server busy")

// flakySender fails the first failures calls and works after that.
type flakySender struct {
	failures int
	calls    int
}

func (f *flakySender) Send(to, body string) error {
	f.calls++
	if f.calls <= f.failures {
		return errBusy
	}
	return nil
}

func TestRetrySucceedsAfterTwoFailures(t *testing.T) {
	sender := &flakySender{failures: 2}

	if err := sendWithRetry(sender, "ann@example.com", "hi", 3); err != nil {
		t.Errorf("sendWithRetry returned %v after %d calls, want success on the third", err, sender.calls)
	}
}

func TestRetryGivesUpAfterTheLastTry(t *testing.T) {
	sender := &flakySender{failures: 10}

	err := sendWithRetry(sender, "ann@example.com", "hi", 3)

	if !errors.Is(err, errBusy) {
		t.Errorf("sendWithRetry = %v, want errBusy", err)
	}
	if sender.calls != 3 {
		t.Errorf("the sender was called %d times, want 3", sender.calls)
	}
}
