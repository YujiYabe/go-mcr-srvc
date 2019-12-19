package recipe

import "app/menu"

// CookDrink ...
type CookDrink struct {
	PrepareDrink PrepareDrink
}

// CookCoffee ...
func (CookDrink *CookDrink) CookCoffee() (users menu.Users, err error) {
	users, err = CookDrink.PrepareDrink.SupplyCoffee()
	return
}

// // UCUIUserByID ...
// func (CookDrink *CookDrink) UCUIUserByID(id int) (user menu.User, err error) {
// 	user, err = CookDrink.PrepareDrink.IFDBFindByID(id)
// 	return
// }

// // UCUIAdd ...
// func (CookDrink *CookDrink) UCUIAdd(u menu.User) (user menu.User, err error) {
// 	user, err = CookDrink.PrepareDrink.IFDBStore(u)
// 	return
// }

// // UCUIUpdate ...
// func (CookDrink *CookDrink) UCUIUpdate(u menu.User) (user menu.User, err error) {
// 	user, err = CookDrink.PrepareDrink.IFDBUpdate(u)
// 	return
// }

// // UCUIDeleteByID ...
// func (CookDrink *CookDrink) UCUIDeleteByID(u menu.User) (err error) {
// 	err = CookDrink.PrepareDrink.IFDBDeleteByID(u)
// 	return
// }
