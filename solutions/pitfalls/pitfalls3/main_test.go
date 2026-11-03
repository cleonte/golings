// pitfalls3
// Make the tests pass!

package main_test

import (
	"strconv"
	"testing"
)

// parseCount reads a number. The empty text is 0. It returns the error of
// strconv.Atoi when the text is not a number.
func parseCount(text string) (count int, err error) {
	if text != "" {
		count, err = strconv.Atoi(text)
		if err != nil {
			return 0, err
		}
	}
	return
}

func TestParseCount(t *testing.T) {
	if got, err := parseCount("42"); got != 42 || err != nil {
		t.Errorf(`parseCount("42") = %d, %v, want 42, nil`, got, err)
	}
	if got, err := parseCount(""); got != 0 || err != nil {
		t.Errorf(`parseCount("") = %d, %v, want 0, nil`, got, err)
	}
	if _, err := parseCount("four"); err == nil {
		t.Error(`parseCount("four") gave no error`)
	}
}
