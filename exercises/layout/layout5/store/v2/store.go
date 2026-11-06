// Package store is the second version of a store, in a folder called v2.
package store

type Store struct{ values map[string]string }

func New() *Store { return &Store{values: map[string]string{}} }

func (s *Store) Set(key, value string) { s.values[key] = value }

func (s *Store) Get(key string) string { return s.values[key] }
