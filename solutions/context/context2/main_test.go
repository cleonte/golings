// context2
// Make the tests pass!

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
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
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
