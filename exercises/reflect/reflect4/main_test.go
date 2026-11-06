// reflect4
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"reflect"
	"testing"
)

type Config struct {
	Name  string
	Ports []int
	Env   map[string]string
}

// sameConfig reports whether two configs have the same content, including the
// slice and the map inside them.
func sameConfig(a, b Config) bool {
	// TODO: Only the name is compared, so configs with other ports or variables count as equal. == does not compile on a struct with a slice or a map; reflect.DeepEqual compares them by content.
	return a.Name == b.Name
}

var _ = reflect.DeepEqual

func TestSameContentIsTheSameConfig(t *testing.T) {
	a := Config{"api", []int{80, 443}, map[string]string{"MODE": "prod"}}
	b := Config{"api", []int{80, 443}, map[string]string{"MODE": "prod"}}

	if !sameConfig(a, b) {
		t.Error("two configs with the same content were reported as different")
	}
}

func TestDifferentContentIsADifferentConfig(t *testing.T) {
	base := Config{"api", []int{80, 443}, map[string]string{"MODE": "prod"}}
	otherPorts := Config{"api", []int{80}, map[string]string{"MODE": "prod"}}
	otherEnv := Config{"api", []int{80, 443}, map[string]string{"MODE": "dev"}}

	if sameConfig(base, otherPorts) {
		t.Error("configs with different ports were reported as the same")
	}
	if sameConfig(base, otherEnv) {
		t.Error("configs with different variables were reported as the same")
	}
}

func TestNilAndEmptyAreNotDeepEqual(t *testing.T) {
	// A known corner of DeepEqual: a nil slice and an empty one have the same
	// length, but they are not deeply equal.
	nilPorts := Config{Name: "api"}
	emptyPorts := Config{Name: "api", Ports: []int{}}

	if sameConfig(nilPorts, emptyPorts) {
		t.Error("a nil slice and an empty slice were reported as the same")
	}
}
