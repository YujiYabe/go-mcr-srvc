package database

import "app/domain"

// IFDBUserRepository ...
type IFDBUserRepository struct {
	IFDBSQLHandler
}

// IFDBFindByID ...
func (IFDBUserRepository *IFDBUserRepository) IFDBFindByID(id int) (user domain.User, err error) {
	if err = IFDBUserRepository.INFRFind(&user, id).Error; err != nil {
		return
	}
	return
}

// IFDBFindAll ...
func (IFDBUserRepository *IFDBUserRepository) IFDBFindAll() (users domain.Users, err error) {
	if err = IFDBUserRepository.INFRFind(&users).Error; err != nil {
		return
	}
	return
}

// IFDBStore ...
func (IFDBUserRepository *IFDBUserRepository) IFDBStore(u domain.User) (user domain.User, err error) {
	if err = IFDBUserRepository.INFRCreate(&u).Error; err != nil {
		return
	}
	user = u
	return
}

// IFDBUpdate ...
func (IFDBUserRepository *IFDBUserRepository) IFDBUpdate(u domain.User) (user domain.User, err error) {
	if err = IFDBUserRepository.INFRSave(&u).Error; err != nil {
		return
	}
	user = u
	return
}

// IFDBDeleteByID ...
func (IFDBUserRepository *IFDBUserRepository) IFDBDeleteByID(user domain.User) (err error) {
	if err = IFDBUserRepository.INFRDelete(&user).Error; err != nil {
		return
	}
	return
}
