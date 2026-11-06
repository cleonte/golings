// http6
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
	"testing/synctest"
	"time"
)

// fetch gets url through the transport rt, and gives up after two seconds.
func fetch(rt http.RoundTripper, url string) (string, error) {
	// TODO: This client waits as long as the server takes. Give it a timeout of two seconds.
	client := &http.Client{Transport: rt}

	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return string(body), err
}

// slowServer is a transport for a server that needs an hour to answer.
type slowServer struct{}

func (slowServer) RoundTrip(req *http.Request) (*http.Response, error) {
	select {
	case <-time.After(time.Hour):
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     http.Header{},
			Body:       io.NopCloser(strings.NewReader("done")),
			Request:    req,
		}, nil
	case <-req.Context().Done():
		return nil, req.Context().Err()
	}
}

func TestFetchGivesUpOnASlowServer(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		start := time.Now()

		_, err := fetch(slowServer{}, "http://example.test/")

		if !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("fetch returned %v, want a timeout", err)
		}
		if elapsed := time.Since(start); elapsed != 2*time.Second {
			t.Errorf("fetch took %v, want 2s", elapsed)
		}
	})
}
