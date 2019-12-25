package usecase

import (
	entity "app/1_entity"
	"app/3_interface/database"
)

// UCUserInteractor ...
type UCUserInteractor struct {
	UCUserRepository UCUserRepository
}

// NewUCUserInteractor ...
func NewUCUserInteractor(SQLHandler database.IFDBSQLHandler) *UCUserInteractor {
	return &UCUserInteractor{UCUserRepository: database.NewIFDBUserRepository(SQLHandler)}
}

// UCUIUserByID ...
func (UCUserInteractor *UCUserInteractor) UCUIUserByID(id int) (user entity.User, err error) {
	user, err = UCUserInteractor.UCUserRepository.IFDBFindByID(id)
	return
}

// UCUIUsers ...
func (UCUserInteractor *UCUserInteractor) UCUIUsers() (users entity.Users, err error) {
	users, err = UCUserInteractor.UCUserRepository.IFDBFindAll()
	return
}

// // UCUIAdd ...
// func (UCUserInteractor *UCUserInteractor) UCUIAdd(u entity.User) (user entity.User, err error) {
// 	user, err = UCUserInteractor.UCUserRepository.IFDBStore(u)
// 	return
// }

// // UCUIUpdate ...
// func (UCUserInteractor *UCUserInteractor) UCUIUpdate(u entity.User) (user entity.User, err error) {
// 	user, err = UCUserInteractor.UCUserRepository.IFDBUpdate(u)
// 	return
// }

// // UCUIDeleteByID ...
// func (UCUserInteractor *UCUserInteractor) UCUIDeleteByID(u entity.User) (err error) {
// 	err = UCUserInteractor.UCUserRepository.IFDBDeleteByID(u)
// 	return
// }
