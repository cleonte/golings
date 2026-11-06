// encoding5
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"encoding/binary"
	"slices"
	"testing"
)

// encode writes n as 4 bytes, the way network protocols do: the biggest byte first.
func encode(n uint32) []byte {
	// TODO: This puts the smallest byte first (little-endian). Network order is the other one, big-endian.
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, n)
	return buf
}

// decode reads 4 bytes written by encode.
func decode(b []byte) uint32 {
	// TODO: Read them in the same order.
	return binary.LittleEndian.Uint32(b)
}

func TestEncode(t *testing.T) {
	got := encode(256)

	if want := []byte{0, 0, 1, 0}; !slices.Equal(got, want) {
		t.Errorf("encode(256) = %v, want %v", got, want)
	}
}

func TestDecode(t *testing.T) {
	if got := decode([]byte{0, 0, 1, 0}); got != 256 {
		t.Errorf("decode of [0 0 1 0] = %d, want 256", got)
	}
}
