package usecase

import (
	database "app/4_database"
)

// UCUserInteractor ...
type UCUserInteractor struct {
	UCUserRepository UCUserRepository
}

// NewUCUserInteractor ...
func NewUCUserInteractor() *UCUserInteractor {
	return &UCUserInteractor{UCUserRepository: database.NewIFDBUserRepository()}
}

// UCUIUserByID ...
func (UCUserInteractor *UCUserInteractor) UCUIUserByID(id int) (user database.User, err error) {
	user, err = UCUserInteractor.UCUserRepository.IFDBFindByID(id)
	return
}

// UCUIUsers ...
func (UCUserInteractor *UCUserInteractor) UCUIUsers() (users database.Users, err error) {
	users, err = UCUserInteractor.UCUserRepository.IFDBFindAll()
	return
}

// // UCUIAdd ...
// func (UCUserInteractor *UCUserInteractor) UCUIAdd(u database.User) (user database.User, err error) {
// 	user, err = UCUserInteractor.UCUserRepository.IFDBStore(u)
// 	return
// }

// // UCUIUpdate ...
// func (UCUserInteractor *UCUserInteractor) UCUIUpdate(u database.User) (user database.User, err error) {
// 	user, err = UCUserInteractor.UCUserRepository.IFDBUpdate(u)
// 	return
// }

// // UCUIDeleteByID ...
// func (UCUserInteractor *UCUserInteractor) UCUIDeleteByID(u database.User) (err error) {
// 	err = UCUserInteractor.UCUserRepository.IFDBDeleteByID(u)
// 	return
// }
