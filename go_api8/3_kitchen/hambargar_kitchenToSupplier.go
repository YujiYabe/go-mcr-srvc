package kitchen

import (
	supplier "app/4_supplier"
)

// HambargarKitchenToSupplier ...
type HambargarKitchenToSupplier interface {
	IFDBFindByName(name string) (supplier.Hambargar, error)
	IFDBFindAll() (supplier.Hambargars, error)
	// IFDBFindByID(id int) (supplier.hambargar, error)
	// IFDBFindAll() (supplier.hambargars, error)
	// IFDBStore(1_supplier.hambargar) (supplier.hambargar, error)
	// IFDBUpdate(1_supplier.hambargar) (supplier.hambargar, error)
	// IFDBDeleteByID(1_supplier.hambargar) error
}
