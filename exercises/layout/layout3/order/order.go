package order

// TODO: The package user imports order, and order imports user: Go does not allow packages that import each other. An order only needs to know the ID of its owner, so keep the ID and drop the import of user.
import "github.com/cleonte/golings/exercises/layout/layout3/user"

type Order struct {
	ID    int
	Owner user.User
}
