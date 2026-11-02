// maps_package2
// Make the tests pass!

package main_test

import (
	"maps"
	"testing"
)

// withOverrides returns the settings with the overrides applied. The defaults
// are shared by the whole program, so they must not change.
func withOverrides(defaults, overrides map[string]string) map[string]string {
	settings := maps.Clone(defaults)
	maps.Copy(settings, overrides)
	return settings
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
