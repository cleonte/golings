package main

// Saver is whatever can keep the counts. It is defined here, where it is used,
// so the code that keeps them (the store package, a fake in a test) does not
// have to know about the Counter.
type Saver interface {
	Save(counts map[string]int) error
}

// Save hands the counts to s.
func (c *Counter) Save(s Saver) error {
	return s.Save(c.counts)
}
