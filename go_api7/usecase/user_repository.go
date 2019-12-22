package usecase

import "app/domain"

// UCUserRepository ...
type UCUserRepository interface {
	IFDBFindByID(id int) (domain.User, error)
	IFDBFindAll() (domain.Users, error)
	IFDBStore(domain.User) (domain.User, error)
	IFDBUpdate(domain.User) (domain.User, error)
	IFDBDeleteByID(domain.User) error
}
