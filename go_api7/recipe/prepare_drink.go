package recipe

import "app/domain"

// PrepareDrink ...
type PrepareDrink interface {
	IFDBFindByID(id int) (domain.User, error)
	IFDBFindAll() (domain.Users, error)
	IFDBStore(domain.User) (domain.User, error)
	IFDBUpdate(domain.User) (domain.User, error)
	IFDBDeleteByID(domain.User) error
}
