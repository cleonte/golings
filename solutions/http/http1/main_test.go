// http1
// Make the tests pass!

package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// created answers a POST: "201 Created", and a line of text.
func created(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "created")
}

func TestCreated(t *testing.T) {
	rec := httptest.NewRecorder()

	created(rec, httptest.NewRequest(http.MethodPost, "/items", nil))

	if rec.Code != http.StatusCreated {
		t.Errorf("status = %d, want %d: the status must be set before the body is written", rec.Code, http.StatusCreated)
	}
	if got := rec.Body.String(); got != "created\n" {
		t.Errorf("body = %q, want %q", got, "created\n")
	}
}
