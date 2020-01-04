package kitchen

import (
	recipe "app/3_kitchen/recipe"
	supplier "app/4_supplier"
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

// // FindAllHambargars ...
// func (HambargarKitchenHandle *HambargarKitchenHandle) FindAllHambargars() (hambargars supplier.StandardHambargars, err error) {
// 	hambargars, err = HambargarKitchenHandle.HambargarKitchenToSupplier.IFDBFindAll()
// 	return
// }

// // FindHambargarByName ...
// func (HambargarKitchenHandle *HambargarKitchenHandle) FindHambargarByName(name string) (hambargar supplier.StandardHambargar, err error) {
// 	hambargar, err = HambargarKitchenHandle.HambargarKitchenToSupplier.IFDBFindByName(name)
// 	return
// }

// Cook ...
// func (HambargarKitchenHandle *HambargarKitchenHandle) Cook(hambargarName string) (hambargar supplier.Hambargar, err error) {
func (HambargarKitchenHandle *HambargarKitchenHandle) Cook(hambargarName string) (hambargar supplier.StandardHambargar, err error) {
	// 材料の取り出し
	// 調理
	// パッキング

	RequestVegetables := map[string]int{}
	// RequestIngredients := []string{}
	if hambargarName == "standard" {
		// RequestVegetables = []string{"tomato", "lettuce"}
		RequestVegetables = map[string]int{"tomato": 1, "lettuce": 2}

		// RequestIngredients = []string{"bans", "cheese"}
	}

	receiveVegetables, err := HambargarKitchenHandle.VegetableKitchenToSupplier.ExtractByNames(RequestVegetables)

	if err != nil {
		return
	}
	hambargar, err = recipe.Combine(receiveVegetables)

	return
}
