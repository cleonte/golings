package main

import "strings"

// Counter counts how many times each word was added, ignoring case.
type Counter struct {
	counts map[string]int
}

func NewCounter() *Counter {
	return &Counter{counts: make(map[string]int)}
}

// Add counts every word of text.
func (c *Counter) Add(text string) {
	for _, word := range strings.Fields(text) {
		c.counts[strings.ToLower(word)]++
	}
}

// Count is how many times word was added.
func (c *Counter) Count(word string) int {
	return c.counts[strings.ToLower(word)]
}

// Total is how many words were added in all, counting repeats.
func (c *Counter) Total() int {
	total := 0
	for _, n := range c.counts {
		total += n
	}
	return total
}
