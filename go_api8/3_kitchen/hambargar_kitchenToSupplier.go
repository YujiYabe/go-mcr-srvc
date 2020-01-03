package kitchen

import (
	supplier "app/4_supplier"
)

// HambargarKitchenToSupplier ...
type HambargarKitchenToSupplier interface {
	IFDBFindByName(name string) (supplier.StandardHambargar, error)
	IFDBFindAll() (supplier.StandardHambargars, error)
	// IFDBFindByID(id int) (supplier.hambargar, error)
	// IFDBFindAll() (supplier.hambargars, error)
	// IFDBStore(1_supplier.hambargar) (supplier.hambargar, error)
	// IFDBUpdate(1_supplier.hambargar) (supplier.hambargar, error)
	// IFDBDeleteByID(1_supplier.hambargar) error
}

// VegetableKitchenToSupplier ...
type VegetableKitchenToSupplier interface {
	ExtractByName(name string) (supplier.Vegetable, error)
	ExtractByNames(RequestVegetables []string) (supplier.Vegetables, error)
}
