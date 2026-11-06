// fakes1
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"errors"
	"fmt"
	"testing"
)

// Sender is anything that can deliver a message.
type Sender interface {
	Send(to, body string) error
}

// sendToTheNetwork stands in for the real mail server.
func sendToTheNetwork(to, body string) error {
	return errors.New("smtp: no network in tests")
}

type Greeter struct {
	Sender Sender
}

// Greet welcomes a new user by message.
func (g Greeter) Greet(name, address string) error {
	// TODO: The Sender field is ignored and the real network is used, so a test can't replace it. Send through g.Sender.
	return sendToTheNetwork(address, fmt.Sprintf("Welcome, %s!", name))
}

// fakeSender remembers what it was asked to send.
type fakeSender struct {
	sent []string
}

func (f *fakeSender) Send(to, body string) error {
	f.sent = append(f.sent, to+": "+body)
	return nil
}

func TestGreetSendsAWelcome(t *testing.T) {
	sender := &fakeSender{}
	greeter := Greeter{Sender: sender}

	if err := greeter.Greet("Ann", "ann@example.com"); err != nil {
		t.Fatalf("Greet returned an error: %v", err)
	}
	want := []string{"ann@example.com: Welcome, Ann!"}
	if len(sender.sent) != 1 || sender.sent[0] != want[0] {
		t.Errorf("sent %q, want %q", sender.sent, want)
	}
}
