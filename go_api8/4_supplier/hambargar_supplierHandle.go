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

// // IFDBFindByName ...
// func (HambargarSupplierHandle *HambargarSupplierHandle) IFDBFindByName(name string) (hambargar StandardHambargar, err error) {
// 	whereParam := new(StandardHambargar)
// 	whereParam.Name = name

// 	if err = HambargarSupplierHandle.HambargarSupplierToStocker.StockFindByName(&hambargar, whereParam).Error; err != nil {
// 		return
// 	}
// 	return
// }

// // IFDBFindAll ...
// func (HambargarSupplierHandle *HambargarSupplierHandle) IFDBFindAll() (hambargars StandardHambargars, err error) {
// 	if err = HambargarSupplierHandle.HambargarSupplierToStocker.StockFind(&hambargars).Error; err != nil {
// 		return
// 	}
// 	return
// }
