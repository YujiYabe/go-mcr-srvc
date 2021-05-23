package database

import (
	entity "app/1_entity"
	"app/4_framework_driver/db"
)

// IFDBUserRepository ...
type IFDBUserRepository struct {
	IFDBSQLHandler
}

// NewIFDBUserRepository ...
func NewIFDBUserRepository() *IFDBUserRepository {
	return &IFDBUserRepository{IFDBSQLHandler: db.NewSQLHandler()}
}

// IFDBFindByID ...
func (IFDBUserRepository *IFDBUserRepository) IFDBFindByID(id int) (user entity.User, err error) {
	if err = IFDBUserRepository.INFRFind(&user, id).Error; err != nil {
		return
	}
	return
}

// IFDBFindAll ...
func (IFDBUserRepository *IFDBUserRepository) IFDBFindAll() (users entity.Users, err error) {
	if err = IFDBUserRepository.INFRFind(&users).Error; err != nil {
		return
	}
	return
}

// // IFDBStore ...
// func (IFDBUserRepository *IFDBUserRepository) IFDBStore(u entity.User) (user entity.User, err error) {
// 	if err = IFDBUserRepository.INFRCreate(&u).Error; err != nil {
// 		return
// 	}
// 	user = u
// 	return
// }

// // IFDBUpdate ...
// func (IFDBUserRepository *IFDBUserRepository) IFDBUpdate(u entity.User) (user entity.User, err error) {
// 	if err = IFDBUserRepository.INFRSave(&u).Error; err != nil {
// 		return
// 	}
// 	user = u
// 	return
// }

// // IFDBDeleteByID ...
// func (IFDBUserRepository *IFDBUserRepository) IFDBDeleteByID(user entity.User) (err error) {
// 	if err = IFDBUserRepository.INFRDelete(&user).Error; err != nil {
// 		return
// 	}
// 	return
// }
