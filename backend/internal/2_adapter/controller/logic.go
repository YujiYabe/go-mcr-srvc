package controller

import (
	"context"

	structObject "backend/internal/4_domain/struct_object"
	valueObject "backend/internal/4_domain/value_object"
)

func (receiver *controller) Start() {}

func (receiver *controller) GetPersonList(
	ctx context.Context,
) (
	personList structObject.PersonList,
	err error,
) {
	personList, err = receiver.UseCase.GetPersonList(ctx)
	return
}

func (receiver *controller) GetPersonByCondition(
	ctx context.Context,
	reqPerson structObject.Person,
) (
	resPersonList structObject.PersonList,
	err error,
) {
	resPersonList, err = receiver.UseCase.GetPersonByCondition(
		ctx,
		reqPerson,
	)
	return
}

func (receiver *controller) FetchAccessToken(
	ctx context.Context,
	credential structObject.Credential,
) (
	accessToken valueObject.AccessToken,
	err error,
) {
	accessToken, err = receiver.UseCase.FetchAccessToken(
		ctx,
		credential,
	)
	return
}

func (receiver *controller) ViaGRPC(
	ctx context.Context,
) (
	err error,
) {
	err = receiver.UseCase.ViaGRPC(
		ctx,
	)
	return
}
