// context4
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"context"
	"errors"
	"testing"
)

var ErrQuotaExceeded = errors.New("quota exceeded")

// limited returns a context and a function that records one use of a resource.
// The context ends with the max-th use.
func limited(parent context.Context, max int) (context.Context, func()) {
	// TODO: A plain cancel says nothing about why the context ended. Create it so that the reason can be given, and give ErrQuotaExceeded when the quota runs out.
	ctx, cancel := context.WithCancel(parent)

	uses := 0
	return ctx, func() {
		uses++
		if uses == max {
			cancel()
		}
	}
}

func TestTheContextEndsWithTheReason(t *testing.T) {
	ctx, use := limited(context.Background(), 2)

	use()
	if err := ctx.Err(); err != nil {
		t.Fatalf("the context ended after one use: %v", err)
	}

	use()
	if err := ctx.Err(); !errors.Is(err, context.Canceled) {
		t.Errorf("ctx.Err() = %v, want context.Canceled", err)
	}
	if cause := context.Cause(ctx); !errors.Is(cause, ErrQuotaExceeded) {
		t.Errorf("context.Cause(ctx) = %v, want %v", cause, ErrQuotaExceeded)
	}
}
