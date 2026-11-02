// maps_package2
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"maps"
	"testing"
)

// withOverrides returns the settings with the overrides applied. The defaults
// are shared by the whole program, so they must not change.
func withOverrides(defaults, overrides map[string]string) map[string]string {
	// TODO: This changes the defaults themselves. Start from a copy of them, and copy the overrides on top.
	for key, value := range overrides {
		defaults[key] = value
	}
	return defaults
}

func TestWithOverrides(t *testing.T) {
	defaults := map[string]string{"theme": "light", "language": "en"}

	got := withOverrides(defaults, map[string]string{"theme": "dark", "font": "mono"})

	if want := (map[string]string{"theme": "dark", "language": "en", "font": "mono"}); !maps.Equal(got, want) {
		t.Errorf("withOverrides gave %v, want %v", got, want)
	}
	if want := (map[string]string{"theme": "light", "language": "en"}); !maps.Equal(defaults, want) {
		t.Errorf("the defaults were changed to %v, want them to stay %v", defaults, want)
	}
}
