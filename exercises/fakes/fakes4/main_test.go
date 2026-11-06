// fakes4
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"testing"
)

// FileStore keeps the names of users in a file on disk.
type FileStore struct {
	Path string
}

func (s *FileStore) Get(id int) (string, error) {
	panic("the tests must not touch the disk")
}

// TODO: report only calls Get, but it asks for a *FileStore, so a test can't give it anything else. Take a small interface that has just the method it uses (define it here, next to the code that uses it) and accept that.
func report(s *FileStore, id int) string {
	name, err := s.Get(id)
	if err != nil {
		return "unknown"
	}
	return "user: " + name
}

type fakeGetter struct{}

func (fakeGetter) Get(id int) (string, error) {
	return "Ann", nil
}

func TestReportUsesTheName(t *testing.T) {
	if got, want := report(fakeGetter{}, 1), "user: Ann"; got != want {
		t.Errorf("report = %q, want %q", got, want)
	}
}
