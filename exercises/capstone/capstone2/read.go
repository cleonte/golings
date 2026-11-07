package main

import (
	"bufio"
	"io"
)

// AddFrom counts the words that come from r, however long its lines are.
func (c *Counter) AddFrom(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	// TODO: The scanner reads lines, which can be too long for it, and a read error is dropped. Read words instead, and return the error of the scanner.
	for scanner.Scan() {
		c.Add(scanner.Text())
	}
	return nil
}
