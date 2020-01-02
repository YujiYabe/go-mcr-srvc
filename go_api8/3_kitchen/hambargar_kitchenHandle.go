package kitchen

import (
	supplier "app/4_supplier"
	"log"
	"runtime"
)

// HambargarKitchenHandle ...
type HambargarKitchenHandle struct {
	HambargarKitchenToSupplier HambargarKitchenToSupplier
	VegetableKitchenToSupplier VegetableKitchenToSupplier
}

// NewHambargarKitchenHandle ...
func NewHambargarKitchenHandle() *HambargarKitchenHandle {
	return &HambargarKitchenHandle{HambargarKitchenToSupplier: supplier.NewHambargarSupplierHandle()}
}

// FindAllHambargars ...
func (HambargarKitchenHandle *HambargarKitchenHandle) FindAllHambargars() (hambargars supplier.Hambargars, err error) {
	hambargars, err = HambargarKitchenHandle.HambargarKitchenToSupplier.IFDBFindAll()
	return
}

// FindHambargarByName ...
func (HambargarKitchenHandle *HambargarKitchenHandle) FindHambargarByName(name string) (hambargar supplier.Hambargar, err error) {
	hambargar, err = HambargarKitchenHandle.HambargarKitchenToSupplier.IFDBFindByName(name)
	return
}

// Cook ...
func (HambargarKitchenHandle *HambargarKitchenHandle) Cook(name string) (hambargar supplier.Hambargar, err error) {
	// hambargar, err = HambargarKitchenHandle.HambargarKitchenToSupplier.IFDBFindByName(name)
	vegetables, err := HambargarKitchenHandle.VegetableKitchenToSupplier.extractByName(name)

	debug := vegetables
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	log.Println("====================================")
	log.Printf("%s:%d %s\n", file, line, f.Name())
	log.Println("====================================")
	log.Printf("%v\n", debug)
	log.Println("------------------------------------")
	log.Printf("%+v\n", debug)
	log.Println("------------------------------------")
	log.Printf("%#v\n", debug)
	log.Println("====================================")

	return
}
