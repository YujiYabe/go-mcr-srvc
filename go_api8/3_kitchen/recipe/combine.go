package recipe

import (
	supplier "app/4_supplier"
	"log"
	"reflect"
)

func Combine(vegetables supplier.Vegetables) (hambargar supplier.StandardHambargar, err error) {

	v := reflect.Indirect(reflect.ValueOf(hambargar))
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		// フィールド名
		println("Field: " + t.Field(i).Name)

		// 値
		f := v.Field(i)
		println("Value: " + f.String())

		// if value, ok := i.(int); ok {
		// 	println("Value: " + strconv.Itoa(value))
		// } else {
		// 	println("Value: " + f.String())
		// }
	}

	debug := vegetables

	log.Println("====================================")
	log.Printf("%v\n", debug)
	log.Println("------------------------------------")
	log.Printf("%+v\n", debug)
	log.Println("------------------------------------")
	log.Printf("%#v\n", debug)
	log.Println("====================================")

	// hambargar, err = HambargarKitchenHandle.HambargarKitchenToSupplier.IFDBFindByName(name)
	return
}
