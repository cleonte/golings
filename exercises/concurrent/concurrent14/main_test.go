// concurrent14
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

// slowLookup takes an hour to answer, unless the context ends first.
func slowLookup(ctx context.Context) (string, error) {
	select {
	case <-time.After(time.Hour):
		return "found", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

// lookupWithin gives slowLookup at most d to answer.
func lookupWithin(d time.Duration) (string, error) {
	// TODO: The lookup takes an hour. Give it a context that ends after d.
	return slowLookup(context.Background())
}

func TestLookupGivesUpAtTheDeadline(t *testing.T) {
	// The clock of the test is fake: an hour passes in no time.
	synctest.Test(t, func(t *testing.T) {
		start := time.Now()

		_, err := lookupWithin(time.Second)

		if !errors.Is(err, context.DeadlineExceeded) {
			t.Errorf("err = %v, want context.DeadlineExceeded", err)
		}
		if elapsed := time.Since(start); elapsed != time.Second {
			t.Errorf("gave up after %v, want after 1s", elapsed)
		}
	})
}

func TestLookupFinishesWhenThereIsEnoughTime(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		got, err := lookupWithin(2 * time.Hour)

		if err != nil || got != "found" {
			t.Errorf("lookupWithin(2h) = %q, %v, want found, nil", got, err)
		}
	})
}
