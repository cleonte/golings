// layout3
// Make me compile!

// I AM NOT DONE
package main

import (
	"fmt"

	"github.com/cleonte/golings/exercises/layout/layout3/order"
	"github.com/cleonte/golings/exercises/layout/layout3/user"
)

func main() {
	ann := user.User{ID: 1, Name: "Ann"}
	o := order.Order{ID: 7, OwnerID: ann.ID}
	fmt.Println(user.Summary(ann, []order.Order{o}))
}
