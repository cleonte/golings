// http4
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

// square answers /square?n=7 with 49.
func square(w http.ResponseWriter, r *http.Request) {
	// TODO: A number that can't be read is ignored, and the answer is 0. Answer such a request with 400 Bad Request (http.Error) and stop.
	n, _ := strconv.Atoi(r.URL.Query().Get("n"))

	fmt.Fprint(w, n*n)
}

func TestSquare(t *testing.T) {
	rec := httptest.NewRecorder()

	square(rec, httptest.NewRequest(http.MethodGet, "/square?n=7", nil))

	if rec.Code != http.StatusOK || rec.Body.String() != "49" {
		t.Errorf("got %d %q, want 200 %q", rec.Code, rec.Body.String(), "49")
	}
}

func TestSquareOfSomethingThatIsNotANumber(t *testing.T) {
	rec := httptest.NewRecorder()

	square(rec, httptest.NewRequest(http.MethodGet, "/square?n=seven", nil))

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
}
