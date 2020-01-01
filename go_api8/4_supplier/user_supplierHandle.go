package supplier

import (
	stocker "app/5_stocker"
)

// UserSupplierHandle ...
type UserSupplierHandle struct {
	SupplierToStocker
}

// NewUserSupplierHandle ...
func NewUserSupplierHandle() *UserSupplierHandle {
	return &UserSupplierHandle{SupplierToStocker: stocker.NewMySQLHandler()}
}

// IFDBFindByID ...
func (UserSupplierHandle *UserSupplierHandle) IFDBFindByID(id int) (user User, err error) {
	if err = UserSupplierHandle.StockFind(&user, id).Error; err != nil {
		return
	}
	return
}

// IFDBFindAll ...
func (UserSupplierHandle *UserSupplierHandle) IFDBFindAll() (users Users, err error) {
	if err = UserSupplierHandle.StockFind(&users).Error; err != nil {
		return
	}
	return
}

// // IFDBStore ...
// func (UserSupplierHandle *UserSupplierHandle) IFDBStore(u User) (user User, err error) {
// 	if err = UserSupplierHandle.INFRCreate(&u).Error; err != nil {
// 		return
// 	}
// 	user = u
// 	return
// }

// // IFDBUpdate ...
// func (UserSupplierHandle *UserSupplierHandle) IFDBUpdate(u User) (user User, err error) {
// 	if err = UserSupplierHandle.INFRSave(&u).Error; err != nil {
// 		return
// 	}
// 	user = u
// 	return
// }

// // IFDBDeleteByID ...
// func (UserSupplierHandle *UserSupplierHandle) IFDBDeleteByID(user User) (err error) {
// 	if err = UserSupplierHandle.INFRDelete(&user).Error; err != nil {
// 		return
// 	}
// 	return
// }
