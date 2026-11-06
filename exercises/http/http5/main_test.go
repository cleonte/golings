// http5
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// upload says how many bytes it received. It accepts at most 10 of them.
func upload(w http.ResponseWriter, r *http.Request) {
	// TODO: Nothing limits the body, so a client can send as much as it likes. Limit it to 10 bytes (the net/http package has a wrapper for the body), and answer 413 Request Entity Too Large when there is more.
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "cannot read the body", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "got %d bytes", len(data))
}

func TestASmallBodyIsAccepted(t *testing.T) {
	rec := httptest.NewRecorder()

	upload(rec, httptest.NewRequest(http.MethodPost, "/upload", strings.NewReader("hello")))

	if rec.Code != http.StatusOK || rec.Body.String() != "got 5 bytes" {
		t.Errorf("got %d %q, want 200 %q", rec.Code, rec.Body.String(), "got 5 bytes")
	}
}

func TestABigBodyIsRefused(t *testing.T) {
	rec := httptest.NewRecorder()

	upload(rec, httptest.NewRequest(http.MethodPost, "/upload", strings.NewReader(strings.Repeat("x", 100))))

	if rec.Code != http.StatusRequestEntityTooLarge {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusRequestEntityTooLarge)
	}
}
