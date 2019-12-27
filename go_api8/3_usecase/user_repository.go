package usecase

import (
	database "app/4_database"
)

// UCUserRepository ...
type UCUserRepository interface {
	IFDBFindByID(id int) (database.User, error)
	IFDBFindAll() (database.Users, error)
	// IFDBStore(1_database.User) (1_database.User, error)
	// IFDBUpdate(1_database.User) (1_database.User, error)
	// IFDBDeleteByID(1_database.User) error
}
