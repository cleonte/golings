// bench3
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"bytes"
	"runtime/pprof"
	"testing"
)

// profileCPU runs work and returns the CPU profile of it, in the pprof format.
func profileCPU(work func()) ([]byte, error) {
	var buf bytes.Buffer
	if err := pprof.StartCPUProfile(&buf); err != nil {
		return nil, err
	}
	work()
	// TODO: The profile is written while it runs and finished when it is stopped, so the buffer is still empty here. Stop the CPU profile before returning.
	return buf.Bytes(), nil
}

func TestProfileHasTheGzipHeader(t *testing.T) {
	total := 0
	data, err := profileCPU(func() {
		for i := range 1_000_000 {
			total += i % 7
		}
	})

	if err != nil {
		t.Fatalf("profileCPU returned an error: %v", err)
	}
	if len(data) < 2 || data[0] != 0x1f || data[1] != 0x8b {
		t.Errorf("the profile has %d bytes and no gzip header: it is not finished", len(data))
	}
	_ = total
}
