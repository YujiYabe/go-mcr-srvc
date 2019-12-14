package usecase

import "app/domain"

// UserInteractor ...
type UCUserInteractor struct {
	UCUserRepository UCUserRepository
}

// UCUIUserByID ...
func (UCuserInteractor *UCUserInteractor) UCUIUserByID(id int) (user domain.User, err error) {
	user, err = UCuserInteractor.UCUserRepository.IFDBFindByID(id)
	return
}

// UCUIUsers ...
func (UCuserInteractor *UCUserInteractor) UCUIUsers() (users domain.Users, err error) {
	users, err = UCuserInteractor.UCUserRepository.IFDBFindAll()
	return
}

// UCUIAdd ...
func (UCuserInteractor *UCUserInteractor) UCUIAdd(u domain.User) (user domain.User, err error) {
	user, err = UCuserInteractor.UCUserRepository.IFDBStore(u)
	return
}

// UCUIUpdate ...
func (UCuserInteractor *UCUserInteractor) UCUIUpdate(u domain.User) (user domain.User, err error) {
	user, err = UCuserInteractor.UCUserRepository.IFDBUpdate(u)
	return
}

// UCUIDeleteByID ...
func (UCuserInteractor *UCUserInteractor) UCUIDeleteByID(u domain.User) (err error) {
	err = UCuserInteractor.UCUserRepository.IFDBDeleteByID(u)
	return
}
