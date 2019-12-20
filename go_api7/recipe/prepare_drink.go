package recipe

import "app/menu"

// PrepareDrink ...
type PrepareDrink interface {
	SupplyCoffee() (menu.Drink, error)
	// IFDBFindByID(id int) (menu.User, error)
	// IFDBStore(menu.User) (menu.User, error)
	// IFDBUpdate(menu.User) (menu.User, error)
	// IFDBDeleteByID(menu.User) error
}
