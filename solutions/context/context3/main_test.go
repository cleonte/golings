// context3
// Make the tests pass!

package main_test

import (
	"context"
	"testing"
)

// A request carries its ID and the ID of its user in the context.

type contextKey int

const (
	requestIDKey contextKey = iota
	userIDKey
)

func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

func RequestID(ctx context.Context) string {
	id, _ := ctx.Value(requestIDKey).(string)
	return id
}

func WithUserID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, userIDKey, id)
}

func UserID(ctx context.Context) string {
	id, _ := ctx.Value(userIDKey).(string)
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
