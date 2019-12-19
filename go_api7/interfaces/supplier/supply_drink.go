package supplier

import "app/menu"

// SupplyDrink ...
type SupplyDrink struct {
	ExtractDrink
}

// SupplyCoffee ...
func (SupplyDrink *SupplyDrink) SupplyCoffee() (users menu.Users, err error) {
	if err = SupplyDrink.extractDrink(&users).Error; err != nil {
		return
	}
	return
}

// // IFDBFindByID ...
// func (SupplyDrink *SupplyDrink) IFDBFindByID(id int) (user menu.User, err error) {
// 	if err = SupplyDrink.extractDrink(&user, id).Error; err != nil {
// 		return
// 	}
// 	return
// }

// // IFDBStore ...
// func (SupplyDrink *SupplyDrink) IFDBStore(u menu.User) (user menu.User, err error) {
// 	if err = SupplyDrink.INFRCreate(&u).Error; err != nil {
// 		return
// 	}
// 	user = u
// 	return
// }

// // IFDBUpdate ...
// func (SupplyDrink *SupplyDrink) IFDBUpdate(u menu.User) (user menu.User, err error) {
// 	if err = SupplyDrink.INFRSave(&u).Error; err != nil {
// 		return
// 	}
// 	user = u
// 	return
// }

// // IFDBDeleteByID ...
// func (SupplyDrink *SupplyDrink) IFDBDeleteByID(user menu.User) (err error) {
// 	if err = SupplyDrink.INFRDelete(&user).Error; err != nil {
// 		return
// 	}
// 	return
// }
