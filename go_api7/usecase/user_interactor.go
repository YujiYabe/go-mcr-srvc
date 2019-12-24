package usecase

import (
	"app/domain"
	"app/interfaces/database"
)

// UCUserInteractor ...
type UCUserInteractor struct {
	UCUserRepository UCUserRepository
}

// NewUCUserInteractor ...
// func NewUCUserInteractor(SQLHandler database.IFDBSQLHandler) UCUserInteractor {
// 	return UCUserInteractor{
// 		UCUserRepository: &database.IFDBUserRepository{
// 			IFDBSQLHandler: SQLHandler,
// 		},
// 	}
// }

// NewUCUserInteractora ...
func NewUCUserInteractora(SQLHandler database.IFDBSQLHandler) *UCUserInteractor {
	return &UCUserInteractor{UCUserRepository: database.NewIFDBUserRepository(SQLHandler)}
}

// UCUIUserByID ...
func (UCUserInteractor *UCUserInteractor) UCUIUserByID(id int) (user domain.User, err error) {
	user, err = UCUserInteractor.UCUserRepository.IFDBFindByID(id)
	return
}

// UCUIUsers ...
func (UCUserInteractor *UCUserInteractor) UCUIUsers() (users domain.Users, err error) {
	users, err = UCUserInteractor.UCUserRepository.IFDBFindAll()
	return
}

// // UCUIAdd ...
// func (UCUserInteractor *UCUserInteractor) UCUIAdd(u domain.User) (user domain.User, err error) {
// 	user, err = UCUserInteractor.UCUserRepository.IFDBStore(u)
// 	return
// }

// // UCUIUpdate ...
// func (UCUserInteractor *UCUserInteractor) UCUIUpdate(u domain.User) (user domain.User, err error) {
// 	user, err = UCUserInteractor.UCUserRepository.IFDBUpdate(u)
// 	return
// }

// // UCUIDeleteByID ...
// func (UCUserInteractor *UCUserInteractor) UCUIDeleteByID(u domain.User) (err error) {
// 	err = UCUserInteractor.UCUserRepository.IFDBDeleteByID(u)
// 	return
// }
