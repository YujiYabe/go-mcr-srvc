package kitchen

import (
	supplier "app/4_supplier"
	"log"
)

// HambargarKitchenHandle ...
type HambargarKitchenHandle struct {
	HambargarKitchenToSupplier HambargarKitchenToSupplier
	VegetableKitchenToSupplier VegetableKitchenToSupplier
}

// NewHambargarKitchenHandle ...
func NewHambargarKitchenHandle() *HambargarKitchenHandle {
	return &HambargarKitchenHandle{
		HambargarKitchenToSupplier: supplier.NewHambargarSupplierHandle(),
		VegetableKitchenToSupplier: supplier.NewVegetableSupplierHandle(),
	}
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
// func (HambargarKitchenHandle *HambargarKitchenHandle) Cook(hambargarName string) (hambargar supplier.Hambargar, err error) {
func (HambargarKitchenHandle *HambargarKitchenHandle) Cook(hambargarName string) (hambargar supplier.Hambargar, err error) {
	// 材料の取り出し
	// 調理
	// パッキング

	RequestVegetables := []string{}
	if hambargarName == "normal" {
		RequestVegetables = []string{"tomato", "lettuce"}
	}
	log.Println("========================= Cook")

	vegetables, err := HambargarKitchenHandle.VegetableKitchenToSupplier.ExtractByNames(RequestVegetables)

	log.Println("====================================")
	log.Printf("%v\n", vegetables)
	log.Println("------------------------------------")
	log.Printf("%+v\n", vegetables)
	log.Println("------------------------------------")
	log.Printf("%#v\n", vegetables)
	log.Println("====================================")
	return
}
