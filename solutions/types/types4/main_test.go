// types4
// Make the tests pass!

package main_test

import (
	"reflect"
	"testing"
)

type UserID int
type OrderID int

func TestIDsAreDifferentTypes(t *testing.T) {
	if reflect.TypeOf(UserID(1)) == reflect.TypeOf(OrderID(1)) {
		t.Error("UserID and OrderID are the same type: a user ID can be used where an order ID is expected")
	}
}

func TestIDsAreNotJustInt(t *testing.T) {
	if reflect.TypeOf(UserID(1)) == reflect.TypeOf(1) {
		t.Error("UserID is the same type as int")
	}
}
