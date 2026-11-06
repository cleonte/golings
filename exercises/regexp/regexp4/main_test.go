// regexp4
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"regexp"
	"slices"
	"testing"
)

var hashtagPattern = regexp.MustCompile(`#\w+`)

// hashtags returns every hashtag in the text, in order.
func hashtags(text string) []string {
	// TODO: FindString stops at the first match. Ask for all of them (the count -1 means no limit).
	if tag := hashtagPattern.FindString(text); tag != "" {
		return []string{tag}
	}
	return nil
}

func TestHashtags(t *testing.T) {
	got := hashtags("Learning #golang with #golings, one #exercise at a time")

	want := []string{"#golang", "#golings", "#exercise"}
	if !slices.Equal(got, want) {
		t.Errorf("hashtags = %q, want %q", got, want)
	}
}

func TestNoHashtags(t *testing.T) {
	if got := hashtags("nothing to see"); len(got) != 0 {
		t.Errorf("hashtags = %q, want none", got)
	}
}
