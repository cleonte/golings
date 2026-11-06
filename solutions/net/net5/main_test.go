// net5
// Make the tests pass!

package main_test

import (
	"bufio"
	"fmt"
	"net"
	"testing"
	"time"
)

// serve answers every client that connects, until the listener is closed.
func serve(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			return
		}
		go handle(conn)
	}
}

// handle reads one line from the client and greets it by that line.
func handle(conn net.Conn) {
	defer conn.Close()

	line, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return
	}
	fmt.Fprintf(conn, "hi %s", line)
}

func TestASilentClientDoesNotBlockTheOthers(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { ln.Close() })
	go serve(ln)

	// The first client connects and says nothing.
	silent, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { silent.Close() })

	fast, err := net.Dial("tcp", ln.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { fast.Close() })
	fast.SetDeadline(time.Now().Add(time.Second))
	fmt.Fprint(fast, "ann\n")

	reply, err := bufio.NewReader(fast).ReadString('\n')

	if err != nil {
		t.Fatal("the second client got no answer while the first one was silent")
	}
	if reply != "hi ann\n" {
		t.Errorf("reply = %q, want %q", reply, "hi ann\n")
	}
}
