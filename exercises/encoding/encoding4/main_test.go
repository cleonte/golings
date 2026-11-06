// encoding4
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"encoding/csv"
	"strings"
	"testing"
)

// writeRecord writes one CSV record, with a new line at the end.
func writeRecord(fields []string) (string, error) {
	var out strings.Builder
	w := csv.NewWriter(&out)

	if err := w.Write(fields); err != nil {
		return "", err
	}
	// TODO: The writer keeps what it is given in a buffer, so nothing has reached out yet. Tell it to write what it holds, and check whether that failed.
	return out.String(), nil
}

func TestWriteRecord(t *testing.T) {
	got, err := writeRecord([]string{"Doe, John", "42"})

	if want := "\"Doe, John\",42\n"; err != nil || got != want {
		t.Errorf("writeRecord = %q, %v, want %q and no error", got, err, want)
	}
}
