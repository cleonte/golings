//go:build go1.27

// go127_3
// Make the tests pass!

package main_test

import "testing"

type Address struct {
	Street string
	City   string
}

type Person struct {
	Address
	Name string
}

// newPerson returns the person with all their details set.
func newPerson(name, street, city string) Person {
	return Person{Name: name, Street: street, City: city}
}

func TestNewPerson(t *testing.T) {
	got := newPerson("Ana", "Rua Augusta 1", "Lisbon")

	if got.Name != "Ana" {
		t.Errorf("Name = %q, want Ana", got.Name)
	}
	if got.Address.Street != "Rua Augusta 1" || got.City != "Lisbon" {
		t.Errorf("the address is %+v, want Rua Augusta 1 in Lisbon", got.Address)
	}
}
