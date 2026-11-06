// cli4
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"flag"
	"io"
	"slices"
	"testing"
)

// parseTags collects the values of a flag that can be repeated: -tag a -tag b.
func parseTags(args []string) ([]string, error) {
	fs := flag.NewFlagSet("tags", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	// TODO: A string flag keeps only the last value it saw. fs.Func gives your function every value, one call per -tag: declare tags first and append each value to it.
	tag := fs.String("tag", "", "a tag; repeat the flag for more than one")
	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	var tags []string
	if *tag != "" {
		tags = append(tags, *tag)
	}
	return tags, nil
}

func TestARepeatedFlagKeepsEveryValue(t *testing.T) {
	got, err := parseTags([]string{"-tag", "go", "-tag", "cli", "-tag", "test"})

	if err != nil {
		t.Fatalf("parseTags returned an error: %v", err)
	}
	if want := []string{"go", "cli", "test"}; !slices.Equal(got, want) {
		t.Errorf("tags = %q, want %q", got, want)
	}
}

func TestNoFlagMeansNoTags(t *testing.T) {
	got, err := parseTags(nil)

	if err != nil || len(got) != 0 {
		t.Errorf("parseTags() = %q, %v, want no tags", got, err)
	}
}
