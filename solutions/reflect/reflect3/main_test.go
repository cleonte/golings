// reflect3
// Make the tests pass!

package main_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

// setName sets the Name field of the struct that p points to.
func setName(p any, name string) error {
	v := reflect.ValueOf(p)
	if v.Kind() != reflect.Pointer || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("%s is not a pointer to a struct", v.Kind())
	}
	field := v.Elem().FieldByName("Name")
	if !field.CanSet() {
		return errors.New("Name cannot be set")
	}
	field.SetString(name)
	return nil
}

type Pet struct {
	Name string
	age  int
}

func TestSetNameThroughAPointer(t *testing.T) {
	pet := Pet{Name: "Rex", age: 3}

	if err := setName(&pet, "Fido"); err != nil {
		t.Fatalf("setName returned an error: %v", err)
	}
	if pet.Name != "Fido" {
		t.Errorf("pet.Name = %q, want %q", pet.Name, "Fido")
	}
}

func TestSetNameNeedsAPointer(t *testing.T) {
	// A struct passed by value is a copy: there is nothing to set, and the
	// function has to say so with an error, not panic.
	if err := setName(Pet{Name: "Rex"}, "Fido"); err == nil {
		t.Error("setName of a struct value returned no error")
	}
}
