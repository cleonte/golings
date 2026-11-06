// fakes1
// Make the tests pass!

package main_test

import (
	"fmt"
	"testing"
)

// Sender is anything that can deliver a message.
type Sender interface {
	Send(to, body string) error
}

type Greeter struct {
	Sender Sender
}

// Greet welcomes a new user by message.
func (g Greeter) Greet(name, address string) error {
	return g.Sender.Send(address, fmt.Sprintf("Welcome, %s!", name))
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
