// http2
// Make the tests pass!

package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// newMux answers GET /items/42 with "item 42", and refuses other methods.
func newMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /items/{id}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "item %s", r.PathValue("id"))
	})
	return mux
}

func TestGetAnItem(t *testing.T) {
	rec := httptest.NewRecorder()

	newMux().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/items/42", nil))

	if rec.Code != http.StatusOK || rec.Body.String() != "item 42" {
		t.Errorf("GET /items/42 gave %d %q, want 200 %q", rec.Code, rec.Body.String(), "item 42")
	}
}

func TestOtherMethodsAreRefused(t *testing.T) {
	rec := httptest.NewRecorder()

	newMux().ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/items/42", nil))

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("POST /items/42 gave %d, want %d", rec.Code, http.StatusMethodNotAllowed)
	}
}
