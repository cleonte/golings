//go:build go1.27

// go127_5
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

// Store is anything that keeps values by name.
type Store interface {
	Get(key string) (any, bool)
}

type MemStore struct {
	data map[string]any
}

// TODO: *MemStore must be a Store. Why isn't it?
func (m *MemStore) Get[T any](key string) (T, bool) {
	value, ok := m.data[key].(T)
	return value, ok
}

func main() {
	var store Store = &MemStore{data: map[string]any{"name": "ana", "age": 31}}

	name, _ := GetAs[string](store, "name")
	age, _ := GetAs[int](store, "age")
	_, isText := GetAs[string](store, "age")

	fmt.Println(name, age, isText)
}
