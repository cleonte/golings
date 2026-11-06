// context4
// Make the tests pass!

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
	ctx, cancel := context.WithCancelCause(parent)

	uses := 0
	return ctx, func() {
		uses++
		if uses == max {
			cancel(ErrQuotaExceeded)
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
