// slog2
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"strings"
	"testing"
)

// newLogger writes JSON records to w, debug records included.
func newLogger(w io.Writer) *slog.Logger {
	// TODO: Use the JSON handler instead of the text one, and let debug records through (the default level is Info).
	return slog.New(slog.NewTextHandler(w, nil))
}

func TestJSONLogger(t *testing.T) {
	var out bytes.Buffer
	logger := newLogger(&out)

	logger.Debug("cache miss", "key", "user:7")
	logger.Info("served", "status", 200)

	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(lines) != 2 {
		t.Fatalf("got %d records, want 2 (the debug one too):\n%s", len(lines), out.String())
	}

	var first map[string]any
	if err := json.Unmarshal([]byte(lines[0]), &first); err != nil {
		t.Fatalf("the first record is not JSON: %v\n%s", err, lines[0])
	}
	if first["level"] != "DEBUG" || first["msg"] != "cache miss" || first["key"] != "user:7" {
		t.Errorf("the first record is %v", first)
	}
}
