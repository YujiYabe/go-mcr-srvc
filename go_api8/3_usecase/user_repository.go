package usecase

import (
	deliver "app/4_deliver"
)

// UCUserRepository ...
type UCUserRepository interface {
	IFDBFindByID(id int) (deliver.User, error)
	IFDBFindAll() (deliver.Users, error)
	// IFDBStore(1_deliver.User) (1_deliver.User, error)
	// IFDBUpdate(1_deliver.User) (1_deliver.User, error)
	// IFDBDeleteByID(1_deliver.User) error
}
