// regexp4
// Make the tests pass!

package main_test

import (
	"regexp"
	"slices"
	"testing"
)

var hashtagPattern = regexp.MustCompile(`#\w+`)

// hashtags returns every hashtag in the text, in order.
func hashtags(text string) []string {
	return hashtagPattern.FindAllString(text, -1)
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
