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

// extractByName ...
func (VegetableSupplierHandle *VegetableSupplierHandle) extractByName(name string) (vegetable Vegetable, err error) {
	whereParam := new(Vegetable)
	whereParam.Name = name

	if err = VegetableSupplierHandle.VegetableSupplierToStocker.StockFindByName(&vegetable, whereParam).Error; err != nil {
		return
	}
	return
}
