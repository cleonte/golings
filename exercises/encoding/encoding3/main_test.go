// encoding3
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"slices"
	"strings"
	"testing"
)

// readRecords splits CSV text into records, each one a list of fields.
func readRecords(text string) ([][]string, error) {
	// TODO: Splitting at every comma breaks a field that has a comma in it, written between quotes. Use the package that reads CSV.
	var records [][]string
	for _, line := range strings.Split(strings.TrimSpace(text), "\n") {
		records = append(records, strings.Split(line, ","))
	}
	return records, nil
}

func TestReadRecords(t *testing.T) {
	text := "name,age\n\"Doe, John\",42\nAda,36\n"

	got, err := readRecords(text)

	want := [][]string{{"name", "age"}, {"Doe, John", "42"}, {"Ada", "36"}}
	if err != nil || !slices.EqualFunc(got, want, slices.Equal) {
		t.Errorf("readRecords = %q, %v, want %q and no error", got, err, want)
	}
}
