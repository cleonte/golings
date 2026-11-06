// net2
// Make the tests pass!

package main_test

import (
	"bufio"
	"io"
	"net"
	"strings"
	"testing"
)

// readMessage reads one message: the text up to the end of the line, without
// the newline.
func readMessage(r io.Reader) (string, error) {
	line, err := bufio.NewReader(r).ReadString('\n')
	return strings.TrimSuffix(line, "\n"), err
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
