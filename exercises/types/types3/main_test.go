// types3
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"testing"
)

type ByteSize int64

// TODO: MB and GB have the same value as KB. Each one is 1024 times the one before: let iota do the shifting (the first line of the block, with the value 0, is not used).
const (
	KB ByteSize = 1 << 10
	MB ByteSize = 1 << 10
	GB ByteSize = 1 << 10
)

func TestSizes(t *testing.T) {
	for _, tc := range []struct {
		name string
		got  ByteSize
		want ByteSize
	}{
		{"KB", KB, 1024},
		{"MB", MB, 1024 * 1024},
		{"GB", GB, 1024 * 1024 * 1024},
	} {
		if tc.got != tc.want {
			t.Errorf("%s = %d, want %d", tc.name, tc.got, tc.want)
		}
	}
}
