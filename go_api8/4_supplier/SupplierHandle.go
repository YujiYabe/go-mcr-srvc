package supplier

import (
	stocker "app/5_stocker"
)

// IFDBUserRepository ...
type IFDBUserRepository struct {
	SupplierToStocker
}

// NewIFDBUserRepository ...
func NewIFDBUserRepository() *IFDBUserRepository {
	return &IFDBUserRepository{SupplierToStocker: stocker.NewMySQLHandler()}
}

// IFDBFindByID ...
func (IFDBUserRepository *IFDBUserRepository) IFDBFindByID(id int) (user User, err error) {
	if err = IFDBUserRepository.StockFind(&user, id).Error; err != nil {
		return
	}
	return
}

// IFDBFindAll ...
func (IFDBUserRepository *IFDBUserRepository) IFDBFindAll() (users Users, err error) {
	if err = IFDBUserRepository.StockFind(&users).Error; err != nil {
		return
	}
	return
}

// // IFDBStore ...
// func (IFDBUserRepository *IFDBUserRepository) IFDBStore(u User) (user User, err error) {
// 	if err = IFDBUserRepository.INFRCreate(&u).Error; err != nil {
// 		return
// 	}
// 	user = u
// 	return
// }

// // IFDBUpdate ...
// func (IFDBUserRepository *IFDBUserRepository) IFDBUpdate(u User) (user User, err error) {
// 	if err = IFDBUserRepository.INFRSave(&u).Error; err != nil {
// 		return
// 	}
// 	user = u
// 	return
// }

// // IFDBDeleteByID ...
// func (IFDBUserRepository *IFDBUserRepository) IFDBDeleteByID(user User) (err error) {
// 	if err = IFDBUserRepository.INFRDelete(&user).Error; err != nil {
// 		return
// 	}
// 	return
// }
