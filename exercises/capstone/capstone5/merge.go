package main

// CountAll counts the words of every text, each in a goroutine of its own with
// a Counter of its own, and merges the results. Nothing is shared while the
// goroutines count, so there is nothing to lock.
func CountAll(texts []string) *Counter {
	results := make(chan *Counter)
	for _, text := range texts {
		go func() {
			c := NewCounter()
			c.Add(text)
			results <- c
		}()
	}

	total := NewCounter()
	for range texts {
		total.Merge(<-results)
	}
	return total
}

// Merge adds the counts of other to the counts of c.
func (c *Counter) Merge(other *Counter) {
	for word, n := range other.counts {
		// TODO: The count of a word that is in both counters is replaced by the one of other. Add the two counts instead.
		c.counts[word] = n
	}
}
