package kitchen

import (
	supplier "app/4_supplier"
)

// HambargarKitchenHandle ...
type HambargarKitchenHandle struct {
	HambargarKitchenToSupplier HambargarKitchenToSupplier
}

// NewHambargarKitchenHandle ...
func NewHambargarKitchenHandle() *HambargarKitchenHandle {
	return &HambargarKitchenHandle{HambargarKitchenToSupplier: supplier.NewHambargarSupplierHandle()}
}

// FindHambargarByName ...
func (HambargarKitchenHandle *HambargarKitchenHandle) FindHambargarByName(name string) (hambargar supplier.Hambargar, err error) {
	hambargar, err = HambargarKitchenHandle.HambargarKitchenToSupplier.IFDBFindByName(name)
	return
}

// FindAllHambargars ...
func (HambargarKitchenHandle *HambargarKitchenHandle) FindAllHambargars() (hambargars supplier.Hambargars, err error) {
	hambargars, err = HambargarKitchenHandle.HambargarKitchenToSupplier.IFDBFindAll()
	return
}

// // UCUIAdd ...
// func (HambargarKitchenHandle *HambargarKitchenHandle) UCUIAdd(u supplier.Hambargar) (hambargar supplier.Hambargar, err error) {
// 	hambargar, err = HambargarKitchenHandle.HambargarKitchenToSupplier.IFDBStore(u)
// 	return
// }

// // UCUIUpdate ...
// func (HambargarKitchenHandle *HambargarKitchenHandle) UCUIUpdate(u supplier.Hambargar) (hambargar supplier.Hambargar, err error) {
// 	hambargar, err = HambargarKitchenHandle.HambargarKitchenToSupplier.IFDBUpdate(u)
// 	return
// }

// // UCUIDeleteByID ...
// func (HambargarKitchenHandle *HambargarKitchenHandle) UCUIDeleteByID(u supplier.Hambargar) (err error) {
// 	err = HambargarKitchenHandle.HambargarKitchenToSupplier.IFDBDeleteByID(u)
// 	return
// }
