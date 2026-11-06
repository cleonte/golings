package user

import (
	"fmt"

	"github.com/cleonte/golings/exercises/layout/layout3/order"
)

type User struct {
	ID   int
	Name string
}

// Summary says how many orders a user has.
func Summary(u User, orders []order.Order) string {
	return fmt.Sprintf("%s has %d order(s)", u.Name, len(orders))
}
