package main

import (
	"bufio"
	"io"
)

// AddFrom counts the words that come from r, however long its lines are.
func (c *Counter) AddFrom(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		c.Add(scanner.Text())
	}
	return scanner.Err()
}
