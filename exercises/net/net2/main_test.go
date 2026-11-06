// net2
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"io"
	"net"
	"testing"
)

// readMessage reads one message: the text up to the end of the line, without
// the newline.
func readMessage(r io.Reader) (string, error) {
	// TODO: One Read returns whatever has arrived, which can be a piece of the message: a connection is a stream, not a list of messages. Read until the newline, for example with a bufio.Reader and ReadString.
	buf := make([]byte, 64)
	n, err := r.Read(buf)
	return string(buf[:n]), err
}

func TestAMessageArrivesInPieces(t *testing.T) {
	client, server := net.Pipe()
	t.Cleanup(func() {
		client.Close()
		server.Close()
	})
	// The sender writes the message in two pieces, as a network may deliver it.
	go func() {
		client.Write([]byte("hel"))
		client.Write([]byte("lo\n"))
	}()

	got, err := readMessage(server)

	if err != nil {
		t.Fatalf("readMessage returned an error: %v", err)
	}
	if got != "hello" {
		t.Errorf("readMessage = %q, want %q", got, "hello")
	}
}
