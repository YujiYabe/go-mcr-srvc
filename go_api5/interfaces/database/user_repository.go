package database

import "app/domain"

// IFDBUserRepository ...
type IFDBUserRepository struct {
	IFDBSQLHandler
}

// IFDBFindByID ...
func (IFDBUserRepository *IFDBUserRepository) IFDBFindByID(id int) (user domain.User, err error) {
	if err = IFDBUserRepository.Find(&user, id).Error; err != nil {
		return
	}
	return
}

// IFDBFindAll ...
func (IFDBUserRepository *IFDBUserRepository) IFDBFindAll() (users domain.Users, err error) {
	if err = IFDBUserRepository.Find(&users).Error; err != nil {
		return
	}
	return
}

// IFDBStore ...
func (IFDBUserRepository *IFDBUserRepository) IFDBStore(u domain.User) (user domain.User, err error) {
	if err = IFDBUserRepository.Create(&u).Error; err != nil {
		return
	}
	user = u
	return
}

// IFDBUpdate ...
func (IFDBUserRepository *IFDBUserRepository) IFDBUpdate(u domain.User) (user domain.User, err error) {
	if err = IFDBUserRepository.Save(&u).Error; err != nil {
		return
	}
	user = u
	return
}

// IFDBDeleteByID ...
func (IFDBUserRepository *IFDBUserRepository) IFDBDeleteByID(user domain.User) (err error) {
	if err = IFDBUserRepository.Delete(&user).Error; err != nil {
		return
	}
	return
}
