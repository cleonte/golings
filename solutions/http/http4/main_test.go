// http4
// Make the tests pass!

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
	n, err := strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		http.Error(w, "n must be a whole number", http.StatusBadRequest)
		return
	}

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
