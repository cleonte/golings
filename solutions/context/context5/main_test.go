// context5
// Make the tests pass!

package main_test

import (
	"context"
	"slices"
	"testing"
)

type userKey struct{}

// audit records what a user did. It refuses to write once ctx is done.
func audit(ctx context.Context, log *[]string, what string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	user, _ := ctx.Value(userKey{}).(string)
	*log = append(*log, user+": "+what)
	return nil
}

// handle serves a request whose client may hang up at any moment. The audit entry has to be written even then.
func handle(ctx context.Context, log *[]string) error {
	return audit(context.WithoutCancel(ctx), log, "deleted the report")
}

func TestTheAuditEntryIsWrittenEvenAfterTheClientLeft(t *testing.T) {
	ctx := context.WithValue(context.Background(), userKey{}, "ana")
	ctx, cancel := context.WithCancel(ctx)
	cancel() // the client hung up

	var log []string
	if err := handle(ctx, &log); err != nil {
		t.Fatalf("handle returned %v, want nil", err)
	}
	if want := []string{"ana: deleted the report"}; !slices.Equal(log, want) {
		t.Errorf("log = %q, want %q", log, want)
	}
}
