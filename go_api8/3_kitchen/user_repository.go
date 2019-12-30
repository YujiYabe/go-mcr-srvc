package kitchen

import (
	supply "app/4_supply"
)

// UCUserRepository ...
type UCUserRepository interface {
	IFDBFindByID(id int) (supply.User, error)
	IFDBFindAll() (supply.Users, error)
	// IFDBStore(1_deliver.User) (supply.User, error)
	// IFDBUpdate(1_deliver.User) (supply.User, error)
	// IFDBDeleteByID(1_deliver.User) error
}
