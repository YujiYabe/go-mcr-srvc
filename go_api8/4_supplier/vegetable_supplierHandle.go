package supplier

import (
	stocker "app/5_stocker"
)

// VegetableSupplierHandle ...
type VegetableSupplierHandle struct {
	VegetableSupplierToStocker VegetableSupplierToStocker
}

// NewVegetableSupplierHandle ...
func NewVegetableSupplierHandle() *VegetableSupplierHandle {
	return &VegetableSupplierHandle{VegetableSupplierToStocker: stocker.NewMySQLHandler()}
}

// ExtractByName ...
func (VegetableSupplierHandle *VegetableSupplierHandle) ExtractByName(name string) (vegetable Vegetable, err error) {
	whereParam := new(Vegetable)
	whereParam.Name = name

	if err = VegetableSupplierHandle.VegetableSupplierToStocker.StockFindByName(&vegetable, whereParam).Error; err != nil {
		return
	}
	return
}

// ExtractByNames ...
func (VegetableSupplierHandle *VegetableSupplierHandle) ExtractByNames(RequestVegetables []string) (vegetables Vegetables, err error) {
	whereParam := RequestVegetables
	if err = VegetableSupplierHandle.VegetableSupplierToStocker.StockFindByNames(&vegetables, whereParam).Error; err != nil {
		return
	}
	return
}
