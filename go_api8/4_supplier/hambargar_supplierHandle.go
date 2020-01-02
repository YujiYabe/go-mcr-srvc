package supplier

import (
	stocker "app/5_stocker"
)

// HambargarSupplierHandle ...
type HambargarSupplierHandle struct {
	HambargarSupplierToStocker HambargarSupplierToStocker
}

// NewHambargarSupplierHandle ...
func NewHambargarSupplierHandle() *HambargarSupplierHandle {
	return &HambargarSupplierHandle{HambargarSupplierToStocker: stocker.NewMySQLHandler()}
}

// IFDBFindByName ...
func (HambargarSupplierHandle *HambargarSupplierHandle) IFDBFindByName(name string) (hambargar Hambargar, err error) {
	whereParam := new(Hambargar)
	whereParam.Name = name

	if err = HambargarSupplierHandle.HambargarSupplierToStocker.StockFindByName(&hambargar, whereParam).Error; err != nil {
		return
	}
	return
}

// IFDBFindAll ...
func (HambargarSupplierHandle *HambargarSupplierHandle) IFDBFindAll() (hambargars Hambargars, err error) {
	if err = HambargarSupplierHandle.HambargarSupplierToStocker.StockFind(&hambargars).Error; err != nil {
		return
	}
	return
}

// // IFDBStore ...
// func (HambargarSupplierHandle *HambargarSupplierHandle) IFDBStore(u Hambargar) (hambargar Hambargar, err error) {
// 	if err = HambargarSupplierHandle.INFRCreate(&u).Error; err != nil {
// 		return
// 	}
// 	hambargar = u
// 	return
// }

// // IFDBUpdate ...
// func (HambargarSupplierHandle *HambargarSupplierHandle) IFDBUpdate(u Hambargar) (hambargar Hambargar, err error) {
// 	if err = HambargarSupplierHandle.INFRSave(&u).Error; err != nil {
// 		return
// 	}
// 	hambargar = u
// 	return
// }

// // IFDBDeleteByID ...
// func (HambargarSupplierHandle *HambargarSupplierHandle) IFDBDeleteByID(hambargar Hambargar) (err error) {
// 	if err = HambargarSupplierHandle.INFRDelete(&hambargar).Error; err != nil {
// 		return
// 	}
// 	return
// }
