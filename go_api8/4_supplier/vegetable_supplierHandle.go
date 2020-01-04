package supplier

import (
	stocker "app/5_stocker"
	"log"
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
func (VegetableSupplierHandle *VegetableSupplierHandle) ExtractByNames(RequestVegetables []string) (vegetables Vegetables, err error) {
	// func (VegetableSupplierHandle *VegetableSupplierHandle) ExtractByNames(RequestVegetables []string) (map[string]string, err error) {
	whereParam := RequestVegetables

	err = VegetableSupplierHandle.VegetableSupplierToStocker.StockFindByNames(&vegetables, whereParam).Error

	if err != nil {
		return
	}
	log.Println("------------------------------------")
	log.Printf("%#v\n", vegetables)
	log.Println("====================================")

	// hoge := Hoge{100, "hello"}
	// m := StructToMap(vegetables)

	// fmt.Println(m)

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
