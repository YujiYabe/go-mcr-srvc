package controller

import (
	"context"

	"backend/internal/4_domain/struct_object"
	"backend/internal/4_domain/value_object"
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

func (receiver *controller) GetAccessToken(
	ctx context.Context,
	credential struct_object.Credential,
) (
	accessToken value_object.AccessToken,
	err error,
) {
	accessToken, err = receiver.UseCase.GetAccessToken(
		ctx,
		credential,
	)

	return
}
