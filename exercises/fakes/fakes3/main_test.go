// fakes3
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"errors"
	"testing"
)

var ErrNotFound = errors.New("not found")

// Store is where the names of users are kept.
type Store interface {
	Get(id int) (string, error)
}

// displayName returns the name of a user, or "guest" when there is no such user.
func displayName(s Store, id int) (string, error) {
	name, err := s.Get(id)
	if errors.Is(err, ErrNotFound) {
		return "guest", nil
	}
	return name, err
}

// fakeStore is an in-memory Store for the tests.
type fakeStore struct {
	names map[int]string
}

func (f fakeStore) Get(id int) (string, error) {
	// TODO: A fake has to behave like the real thing, errors included. The real store reports a missing user with ErrNotFound, and this one returns an empty name and no error.
	return f.names[id], nil
}

func TestKnownUserKeepsTheName(t *testing.T) {
	store := fakeStore{names: map[int]string{1: "Ann"}}

	name, err := displayName(store, 1)

	if err != nil || name != "Ann" {
		t.Errorf("displayName(1) = %q, %v, want %q", name, err, "Ann")
	}
}

func TestMissingUserIsAGuest(t *testing.T) {
	store := fakeStore{names: map[int]string{1: "Ann"}}

	name, err := displayName(store, 2)

	if err != nil || name != "guest" {
		t.Errorf("displayName(2) = %q, %v, want %q", name, err, "guest")
	}
}
