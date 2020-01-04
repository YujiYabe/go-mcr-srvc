package recipe

import (
	supplier "app/4_supplier"
	// "reflect"
)

// Combine ...
func Combine(vegetables map[string]int) (hambargar supplier.StandardHambargar, err error) {
	hambargar.Tomato = vegetables["tomato"]

	// v := reflect.Indirect(reflect.ValueOf(hambargar))

	// t := v.Type()

	// for i := 0; i < t.NumField(); i++ {
	// 	// フィールド名
	// 	println("Field: " + t.Field(i).Name)

	// 	// 値
	// 	f := v.Field(i)
	// 	println("Value: " + f.String())

	// }

	// hambargar, err = HambargarKitchenHandle.HambargarKitchenToSupplier.IFDBFindByName(name)
	return
}
