package main

// TODO: The counts have to go to something that can keep them, and the test passes a fake and a text store. Define a Saver interface here (with the one method that Counter needs, Save(counts map[string]int) error), take it instead of any, and call it.

// Save hands the counts to s.
func (c *Counter) Save(s any) error {
	return nil
}
