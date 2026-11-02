// json1
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"encoding/json"
	"testing"
)

// User is sent to the clients as JSON: lower case names, no email when there
// is none, and never the password.
type User struct {
	// TODO: Add tags to the fields: lower case names, omitempty for the email, and "-" for the password.
	ID       int
	Name     string
	Email    string
	Password string
}

func TestUserJSON(t *testing.T) {
	user := User{ID: 7, Name: "ana", Password: "s3cret"}

	data, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	want := `{"id":7,"name":"ana"}`
	if string(data) != want {
		t.Errorf("got %s, want %s", data, want)
	}

	data, err = json.Marshal(User{ID: 8, Name: "carl", Email: "carl@example.com"})
	if err != nil {
		t.Fatal(err)
	}
	want = `{"id":8,"name":"carl","email":"carl@example.com"}`
	if string(data) != want {
		t.Errorf("got %s, want %s", data, want)
	}
}
