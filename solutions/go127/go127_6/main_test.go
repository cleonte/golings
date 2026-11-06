//go:build go1.27

// go127_6
// Make the tests pass!

package main_test

import (
	"encoding/json/v2"
	"testing"
)

type Profile struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

func TestAnEmptySliceIsAnEmptyArray(t *testing.T) {
	data, err := json.Marshal(Profile{Name: "ana"})
	if err != nil {
		t.Fatal(err)
	}

	want := `{"name":"ana","tags":[]}`
	if string(data) != want {
		t.Errorf("got %s, want %s", data, want)
	}
}

func TestADuplicateNameIsAnError(t *testing.T) {
	var profile Profile

	err := json.Unmarshal([]byte(`{"name":"ana","name":"carl"}`), &profile)

	if err == nil {
		t.Errorf("Unmarshal worked and gave the name %q, want an error for the repeated name", profile.Name)
	}
}
