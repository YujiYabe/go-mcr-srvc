package kitchen

import (
	supplier "app/4_supplier"
)

// UserKitchenHandle ...
type UserKitchenHandle struct {
	UserKitchenToSupplier UserKitchenToSupplier
}

// NewUserKitchenHandle ...
func NewUserKitchenHandle() *UserKitchenHandle {
	return &UserKitchenHandle{UserKitchenToSupplier: supplier.NewUserSupplierHandle()}
}

// FindUserByID ...
func (UserKitchenHandle *UserKitchenHandle) FindUserByID(id int) (user supplier.User, err error) {
	user, err = UserKitchenHandle.UserKitchenToSupplier.IFDBFindByID(id)
	return
}

// FindAllUsers ...
func (UserKitchenHandle *UserKitchenHandle) FindAllUsers() (users supplier.Users, err error) {
	users, err = UserKitchenHandle.UserKitchenToSupplier.IFDBFindAll()
	return
}

// // UCUIAdd ...
// func (UserKitchenHandle *UserKitchenHandle) UCUIAdd(u supplier.User) (user supplier.User, err error) {
// 	user, err = UserKitchenHandle.UserKitchenToSupplier.IFDBStore(u)
// 	return
// }

// // UCUIUpdate ...
// func (UserKitchenHandle *UserKitchenHandle) UCUIUpdate(u supplier.User) (user supplier.User, err error) {
// 	user, err = UserKitchenHandle.UserKitchenToSupplier.IFDBUpdate(u)
// 	return
// }

// // UCUIDeleteByID ...
// func (UserKitchenHandle *UserKitchenHandle) UCUIDeleteByID(u supplier.User) (err error) {
// 	err = UserKitchenHandle.UserKitchenToSupplier.IFDBDeleteByID(u)
// 	return
// }
