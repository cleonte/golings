// memory1
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"strings"
	"testing"
	"unique"
)

// intern returns the canonical handle for s: equal strings give equal handles,
// and comparing two handles is as cheap as comparing two pointers.
func intern(s string) unique.Handle[string] {
	// TODO: This always gives the handle of the empty string. Make the handle of s with the unique package.
	return unique.Make("")
}

func TestInternKeepsTheValue(t *testing.T) {
	if got := intern("gopher").Value(); got != "gopher" {
		t.Errorf("intern(\"gopher\").Value() = %q, want %q", got, "gopher")
	}
}

func TestEqualStringsGiveEqualHandles(t *testing.T) {
	a := intern("gopher")
	b := intern(strings.Clone("gopher")) // another string, with the same content

	if a != b {
		t.Error("two equal strings gave handles that are not equal")
	}
	if a == intern("gophers") {
		t.Error("two different strings gave equal handles")
	}
}
