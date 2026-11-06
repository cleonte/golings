// context2
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"context"
	"errors"
	"testing"
	"testing/synctest"
	"time"
)

// slowLookup answers after five seconds, unless ctx is done first.
func slowLookup(ctx context.Context) (string, error) {
	select {
	case <-time.After(5 * time.Second):
		return "found", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

// lookup may spend one second on slowLookup, no more.
func lookup(ctx context.Context) (string, error) {
	// TODO: slowLookup gets as much time as it wants. Derive a context that ends after one second, and release it when lookup returns.
	return slowLookup(ctx)
}

func TestLookupGivesUpAfterASecond(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		start := time.Now()

		_, err := lookup(context.Background())

		if !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("lookup returned %v, want context.DeadlineExceeded", err)
		}
		if elapsed := time.Since(start); elapsed != time.Second {
			t.Errorf("lookup took %v, want 1s", elapsed)
		}
	})
}
