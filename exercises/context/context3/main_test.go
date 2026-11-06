// context3
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"context"
	"testing"
)

// A request carries its ID and the ID of its user in the context.

// TODO: Both values are stored under the same string key, "id", so the second hides the first (and so could any other package that picks "id"). Give each value a key of its own: a type that only this package can name.
func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, "id", id)
}

func RequestID(ctx context.Context) string {
	id, _ := ctx.Value("id").(string)
	return id
}

func WithUserID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, "id", id)
}

func UserID(ctx context.Context) string {
	id, _ := ctx.Value("id").(string)
	return id
}

func TestTheTwoIDsDoNotCollide(t *testing.T) {
	ctx := WithUserID(WithRequestID(context.Background(), "req-7"), "ana")

	if got := RequestID(ctx); got != "req-7" {
		t.Errorf("RequestID = %q, want %q", got, "req-7")
	}
	if got := UserID(ctx); got != "ana" {
		t.Errorf("UserID = %q, want %q", got, "ana")
	}
}

func TestAnIDThatWasNeverSetIsEmpty(t *testing.T) {
	if got := RequestID(context.Background()); got != "" {
		t.Errorf("RequestID of an empty context = %q, want an empty string", got)
	}
}
