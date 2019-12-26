package usecase

import (
	entity "app/1_entity"
)

// UCUserRepository ...
type UCUserRepository interface {
	IFDBFindByID(id int) (entity.User, error)
	IFDBFindAll() (entity.Users, error)
	// IFDBStore(1_entity.User) (1_entity.User, error)
	// IFDBUpdate(1_entity.User) (1_entity.User, error)
	// IFDBDeleteByID(1_entity.User) error
}
