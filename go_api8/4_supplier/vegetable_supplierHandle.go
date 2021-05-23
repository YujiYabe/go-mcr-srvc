package supplier

import (
	stocker "app/5_stocker"
	"reflect"
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
// func (VegetableSupplierHandle *VegetableSupplierHandle) ExtractByNames(RequestVegetables []string) (map[string]string, err error) {
func (VegetableSupplierHandle *VegetableSupplierHandle) ExtractByNames(RequestVegetables map[string]int) (receiveVegetables map[string]int, err error) {
	whereParam := RequestVegetables

	// err = VegetableSupplierHandle.VegetableSupplierToStocker.StockFindByNames(&vegetables, whereParam).Error
	err = VegetableSupplierHandle.VegetableSupplierToStocker.StockFindByNames(whereParam).Error

	if err != nil {
		return
	}
	receiveVegetables = RequestVegetables

	return
}

// StructToMap ...
func StructToMap(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := elem.Type().Field(i).Name
		value := elem.Field(i).Interface()
		result[field] = value
	}

	return result
}
