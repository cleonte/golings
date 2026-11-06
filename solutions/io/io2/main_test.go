// io2
// Make the tests pass!

package main_test

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

// report writes the name and the quantity of each item, one per line, to w.
func report(w io.Writer, items map[string]int, order []string) {
	for _, name := range order {
		fmt.Fprintf(w, "%s: %d\n", name, items[name])
	}
}

func TestReportWritesToTheWriter(t *testing.T) {
	var out strings.Builder

	report(&out, map[string]int{"apples": 3, "pears": 5}, []string{"apples", "pears"})

	if want := "apples: 3\npears: 5\n"; out.String() != want {
		t.Errorf("report wrote %q, want %q", out.String(), want)
	}
}
