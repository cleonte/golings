// Package store keeps the counts of the word counter.
package store

import (
	"fmt"
	"io"
	"maps"
	"slices"
)

// Text writes the counts as lines of "word count", sorted by word.
type Text struct {
	w io.Writer
}

func NewText(w io.Writer) *Text {
	return &Text{w: w}
}

func (t *Text) Save(counts map[string]int) error {
	for _, word := range slices.Sorted(maps.Keys(counts)) {
		if _, err := fmt.Fprintf(t.w, "%s %d\n", word, counts[word]); err != nil {
			return err
		}
	}
	return nil
}
