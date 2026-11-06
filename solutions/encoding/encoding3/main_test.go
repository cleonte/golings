// encoding3
// Make the tests pass!

package main_test

import (
	"encoding/csv"
	"slices"
	"strings"
	"testing"
)

// readRecords splits CSV text into records, each one a list of fields.
func readRecords(text string) ([][]string, error) {
	return csv.NewReader(strings.NewReader(text)).ReadAll()
}

func TestReadRecords(t *testing.T) {
	text := "name,age\n\"Doe, John\",42\nAda,36\n"

	got, err := readRecords(text)

	want := [][]string{{"name", "age"}, {"Doe, John", "42"}, {"Ada", "36"}}
	if err != nil || !slices.EqualFunc(got, want, slices.Equal) {
		t.Errorf("readRecords = %q, %v, want %q and no error", got, err, want)
	}
}
