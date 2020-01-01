package kitchen

import (
	supplier "app/4_supplier"
)

// KitchenToSupplier ...
type KitchenToSupplier interface {
	IFDBFindByID(id int) (supplier.User, error)
	IFDBFindAll() (supplier.Users, error)
	// IFDBStore(1_supplier.User) (supplier.User, error)
	// IFDBUpdate(1_supplier.User) (supplier.User, error)
	// IFDBDeleteByID(1_supplier.User) error
}
