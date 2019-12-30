package kitchen

import (
	supplier "app/4_supplier"
)

// UCUserInteractor ...
type UCUserInteractor struct {
	UCUserRepository UCUserRepository
}

// NewUCUserInteractor ...
func NewUCUserInteractor() *UCUserInteractor {
	return &UCUserInteractor{UCUserRepository: supplier.NewIFDBUserRepository()}
}

// UCUIUserByID ...
func (UCUserInteractor *UCUserInteractor) UCUIUserByID(id int) (user supplier.User, err error) {
	user, err = UCUserInteractor.UCUserRepository.IFDBFindByID(id)
	return
}

// UCUIUsers ...
func (UCUserInteractor *UCUserInteractor) UCUIUsers() (users supplier.Users, err error) {
	users, err = UCUserInteractor.UCUserRepository.IFDBFindAll()
	return
}

// // UCUIAdd ...
// func (UCUserInteractor *UCUserInteractor) UCUIAdd(u supplier.User) (user supplier.User, err error) {
// 	user, err = UCUserInteractor.UCUserRepository.IFDBStore(u)
// 	return
// }

// // UCUIUpdate ...
// func (UCUserInteractor *UCUserInteractor) UCUIUpdate(u supplier.User) (user supplier.User, err error) {
// 	user, err = UCUserInteractor.UCUserRepository.IFDBUpdate(u)
// 	return
// }

// // UCUIDeleteByID ...
// func (UCUserInteractor *UCUserInteractor) UCUIDeleteByID(u supplier.User) (err error) {
// 	err = UCUserInteractor.UCUserRepository.IFDBDeleteByID(u)
// 	return
// }
