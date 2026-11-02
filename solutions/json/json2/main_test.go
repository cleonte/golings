// json2
// Make the tests pass!

package main_test

import (
	"encoding/json"
	"testing"
	"time"
)

// Task has a time that is only set once the task is done.
type Task struct {
	Title  string    `json:"title"`
	DoneAt time.Time `json:"done_at,omitzero"`
}

func TestTaskWithoutATime(t *testing.T) {
	data, err := json.Marshal(Task{Title: "write tests"})
	if err != nil {
		t.Fatal(err)
	}

	want := `{"title":"write tests"}`
	if string(data) != want {
		t.Errorf("got %s, want %s", data, want)
	}
}

func TestTaskWithATime(t *testing.T) {
	done := time.Date(2026, time.February, 1, 10, 30, 0, 0, time.UTC)

	data, err := json.Marshal(Task{Title: "ship it", DoneAt: done})
	if err != nil {
		t.Fatal(err)
	}

	want := `{"title":"ship it","done_at":"2026-02-01T10:30:00Z"}`
	if string(data) != want {
		t.Errorf("got %s, want %s", data, want)
	}
}
