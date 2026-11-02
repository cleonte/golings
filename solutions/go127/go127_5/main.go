//go:build go1.27

// go127_5
// Make me compile!

package main

import "fmt"

// Store is anything that keeps values by name.
type Store interface {
	Get(key string) (any, bool)
}

type MemStore struct {
	data map[string]any
}

func (m *MemStore) Get(key string) (any, bool) {
	value, ok := m.data[key]
	return value, ok
}

// GetAs gives the value as the type T, when it is one.
func GetAs[T any](s Store, key string) (T, bool) {
	value, ok := s.Get(key)
	typed, isT := value.(T)
	return typed, ok && isT
}

func main() {
	var store Store = &MemStore{data: map[string]any{"name": "ana", "age": 31}}

	name, _ := GetAs[string](store, "name")
	age, _ := GetAs[int](store, "age")
	_, isText := GetAs[string](store, "age")

	fmt.Println(name, age, isText)
}
