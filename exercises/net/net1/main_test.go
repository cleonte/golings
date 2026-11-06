// net1
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"bufio"
	"net"
	"testing"
	"time"
)

// handle answers a line that says PING with a line that says PONG, then hangs up.
func handle(conn net.Conn) {
	defer conn.Close()

	line, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return
	}
	if line == "PING\n" {
		// TODO: The server reads the request and never answers. Write the reply "PONG\n" to conn (a net.Conn is an io.Writer too).
	}
}

func TestPingGetsPong(t *testing.T) {
	// net.Pipe is a connection in memory: whatever one end writes, the other reads.
	client, server := net.Pipe()
	t.Cleanup(func() { client.Close() })
	go handle(server)

	client.SetDeadline(time.Now().Add(time.Second))
	if _, err := client.Write([]byte("PING\n")); err != nil {
		t.Fatalf("writing the request: %v", err)
	}
	reply, err := bufio.NewReader(client).ReadString('\n')

	if err != nil {
		t.Fatal("the server sent no reply to PING")
	}
	if reply != "PONG\n" {
		t.Errorf("reply = %q, want %q", reply, "PONG\n")
	}
}
