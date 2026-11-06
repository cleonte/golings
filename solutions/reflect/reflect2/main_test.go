// reflect2
// Make the tests pass!

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
		names = append(names, field.Tag.Get("json"))
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
