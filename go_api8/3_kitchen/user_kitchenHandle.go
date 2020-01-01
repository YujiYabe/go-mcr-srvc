package kitchen

import (
	supplier "app/4_supplier"
)

// UserKitchenHandle ...
type UserKitchenHandle struct {
	KitchenToSupplier KitchenToSupplier
}

// NewUserKitchenHandle ...
func NewUserKitchenHandle() *UserKitchenHandle {
	return &UserKitchenHandle{KitchenToSupplier: supplier.NewIFDBUserRepository()}
}

// UCUIUserByID ...
func (UserKitchenHandle *UserKitchenHandle) UCUIUserByID(id int) (user supplier.User, err error) {
	user, err = UserKitchenHandle.KitchenToSupplier.IFDBFindByID(id)
	return
}

// UCUIUsers ...
func (UserKitchenHandle *UserKitchenHandle) UCUIUsers() (users supplier.Users, err error) {
	users, err = UserKitchenHandle.KitchenToSupplier.IFDBFindAll()
	return
}

// // UCUIAdd ...
// func (UserKitchenHandle *UserKitchenHandle) UCUIAdd(u supplier.User) (user supplier.User, err error) {
// 	user, err = UserKitchenHandle.KitchenToSupplier.IFDBStore(u)
// 	return
// }

// // UCUIUpdate ...
// func (UserKitchenHandle *UserKitchenHandle) UCUIUpdate(u supplier.User) (user supplier.User, err error) {
// 	user, err = UserKitchenHandle.KitchenToSupplier.IFDBUpdate(u)
// 	return
// }

// // UCUIDeleteByID ...
// func (UserKitchenHandle *UserKitchenHandle) UCUIDeleteByID(u supplier.User) (err error) {
// 	err = UserKitchenHandle.KitchenToSupplier.IFDBDeleteByID(u)
// 	return
// }
