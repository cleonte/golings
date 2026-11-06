// net3
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"errors"
	"io"
	"net"
	"os"
	"testing"
	"time"
)

// greet sends a greeting to the first client that connects and is done.
func greet(ln net.Listener) {
	conn, err := ln.Accept()
	if err != nil {
		return
	}
	// TODO: The connection is never closed, so a client that reads until the end of the data waits forever. Close the connection when the greeting is sent.
	conn.Write([]byte("hello\n"))
}

func TestGreetingEndsWithTheConnection(t *testing.T) {
	// Port 0 asks the system for a free port; the listener knows which one.
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { ln.Close() })
	go greet(ln)

	conn, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { conn.Close() })

	conn.SetDeadline(time.Now().Add(time.Second))
	data, err := io.ReadAll(conn)

	if errors.Is(err, os.ErrDeadlineExceeded) {
		t.Fatal("still waiting for the end of the data after a second: the server did not close the connection")
	}
	if err != nil {
		t.Fatalf("reading the greeting: %v", err)
	}
	if string(data) != "hello\n" {
		t.Errorf("greeting = %q, want %q", data, "hello\n")
	}
}
