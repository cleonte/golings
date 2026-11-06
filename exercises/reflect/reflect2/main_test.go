// reflect2
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"reflect"
	"slices"
	"testing"
)

// jsonNames returns the names the json package would use for the fields of a
// struct: the name in the `json:"..."` tag of each field, in order.
func jsonNames(v any) []string {
	var names []string
	for field := range reflect.TypeOf(v).Fields() {
		// TODO: This collects the names of the Go fields. The tag is a string in the struct field: field.Tag.Get("json") returns what is written under the key json.
		names = append(names, field.Name)
	}
	return names
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email_address"`
}

func TestJSONNamesComeFromTheTags(t *testing.T) {
	got := jsonNames(User{})
	want := []string{"id", "name", "email_address"}

	if !slices.Equal(got, want) {
		t.Errorf("jsonNames(User{}) = %q, want %q", got, want)
	}
}
