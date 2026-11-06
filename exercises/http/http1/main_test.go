// http1
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// created answers a POST: "201 Created", and a line of text.
func created(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "created")
	// TODO: The status line goes out with the first byte of the body, so by now it is 200 and this call comes too late. Set the status before writing the body.
	w.WriteHeader(http.StatusCreated)
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
