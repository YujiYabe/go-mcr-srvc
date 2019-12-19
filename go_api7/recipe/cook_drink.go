package recipe

import "app/domain"

// CookDrink ...
type CookDrink struct {
	PrepareDrink PrepareDrink
}

// UCUIUserByID ...
func (CookDrink *CookDrink) UCUIUserByID(id int) (user domain.User, err error) {
	user, err = CookDrink.PrepareDrink.IFDBFindByID(id)
	return
}

// UCUIUsers ...
func (CookDrink *CookDrink) UCUIUsers() (users domain.Users, err error) {
	users, err = CookDrink.PrepareDrink.IFDBFindAll()
	return
}

// UCUIAdd ...
func (CookDrink *CookDrink) UCUIAdd(u domain.User) (user domain.User, err error) {
	user, err = CookDrink.PrepareDrink.IFDBStore(u)
	return
}

// UCUIUpdate ...
func (CookDrink *CookDrink) UCUIUpdate(u domain.User) (user domain.User, err error) {
	user, err = CookDrink.PrepareDrink.IFDBUpdate(u)
	return
}

// UCUIDeleteByID ...
func (CookDrink *CookDrink) UCUIDeleteByID(u domain.User) (err error) {
	err = CookDrink.PrepareDrink.IFDBDeleteByID(u)
	return
}
