// types3
// Make the tests pass!

package main_test

import (
	"testing"
)

type ByteSize int64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
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
