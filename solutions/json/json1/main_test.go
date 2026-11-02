// json1
// Make the tests pass!

package main_test

import (
	"encoding/json"
	"testing"
)

// User is sent to the clients as JSON: lower case names, no email when there
// is none, and never the password.
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	Password string `json:"-"`
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
