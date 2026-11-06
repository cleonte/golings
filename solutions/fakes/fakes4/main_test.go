// fakes4
// Make the tests pass!

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

// getter is the one thing report needs from a store.
type getter interface {
	Get(id int) (string, error)
}

func report(s getter, id int) string {
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
