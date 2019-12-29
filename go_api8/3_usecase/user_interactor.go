package usecase

import (
	deliver "app/4_deliver"
)

// UCUserInteractor ...
type UCUserInteractor struct {
	UCUserRepository UCUserRepository
}

// NewUCUserInteractor ...
func NewUCUserInteractor() *UCUserInteractor {
	return &UCUserInteractor{UCUserRepository: deliver.NewIFDBUserRepository()}
}

// UCUIUserByID ...
func (UCUserInteractor *UCUserInteractor) UCUIUserByID(id int) (user deliver.User, err error) {
	user, err = UCUserInteractor.UCUserRepository.IFDBFindByID(id)
	return
}

// UCUIUsers ...
func (UCUserInteractor *UCUserInteractor) UCUIUsers() (users deliver.Users, err error) {
	users, err = UCUserInteractor.UCUserRepository.IFDBFindAll()
	return
}

// // UCUIAdd ...
// func (UCUserInteractor *UCUserInteractor) UCUIAdd(u deliver.User) (user deliver.User, err error) {
// 	user, err = UCUserInteractor.UCUserRepository.IFDBStore(u)
// 	return
// }

// // UCUIUpdate ...
// func (UCUserInteractor *UCUserInteractor) UCUIUpdate(u deliver.User) (user deliver.User, err error) {
// 	user, err = UCUserInteractor.UCUserRepository.IFDBUpdate(u)
// 	return
// }

// // UCUIDeleteByID ...
// func (UCUserInteractor *UCUserInteractor) UCUIDeleteByID(u deliver.User) (err error) {
// 	err = UCUserInteractor.UCUserRepository.IFDBDeleteByID(u)
// 	return
// }
