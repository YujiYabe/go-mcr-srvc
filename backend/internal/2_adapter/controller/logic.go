package controller

import (
	"context"

	"backend/internal/4_domain/struct_object"
)

func (receiver *controller) Start() {}

func (receiver *controller) GetPersonList(
	ctx context.Context,
) (
	personList struct_object.PersonList,
	err error,
) {

	return receiver.UseCase.GetPersonList(ctx)
}

func (receiver *controller) GetPersonByCondition(
	ctx context.Context,
	reqPerson struct_object.Person,
) (
	resPersonList struct_object.PersonList,
	err error,
) {
	resPersonList, err = receiver.UseCase.GetPersonByCondition(
		ctx,
		reqPerson,
	)

	return resPersonList, err
}
