// http3
// Make the tests pass!

package main_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// requireToken lets a request through to next only when it carries the right bearer token.
func requireToken(token string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer "+token {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func secret() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("secret data"))
	})
}

func TestWithoutTheTokenNothingGetsThrough(t *testing.T) {
	rec := httptest.NewRecorder()

	requireToken("s3cret", secret()).ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))

	if rec.Code != http.StatusUnauthorized {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusUnauthorized)
	}
	if strings.Contains(rec.Body.String(), "secret data") {
		t.Errorf("the protected handler ran anyway: body = %q", rec.Body.String())
	}
}

func TestWithTheTokenTheRequestGoesThrough(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", "Bearer s3cret")

	requireToken("s3cret", secret()).ServeHTTP(rec, req)

	if rec.Code != http.StatusOK || rec.Body.String() != "secret data" {
		t.Errorf("got %d %q, want 200 %q", rec.Code, rec.Body.String(), "secret data")
	}
}
