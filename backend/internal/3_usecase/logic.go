package usecase

import (
	"context"

	"backend/internal/4_domain/struct_object"
	"backend/internal/4_domain/value_object"
)

// Start ...
func (receiver *useCase) Start() {
}

func (receiver *useCase) GetPersonList(
	ctx context.Context,
) (
	personList struct_object.PersonList,
	err error,
) {
	personList, err = receiver.ToGateway.GetPersonList(ctx)
	return
}

func (receiver *useCase) GetPersonByCondition(
	ctx context.Context,
	reqPerson struct_object.Person,
) (
	resPersonList struct_object.PersonList,
	err error,
) {
	resPersonList, err = receiver.ToGateway.GetPersonByCondition(
		ctx,
		reqPerson,
	)
	return
}

func (receiver *useCase) FetchAccessToken(
	ctx context.Context,
	credential struct_object.Credential,
) (
	accessToken value_object.AccessToken,
	err error,
) {
	accessToken, err = receiver.ToGateway.FetchAccessToken(
		ctx,
		credential,
	)
	return
}

func (receiver *useCase) ViaGRPC(
	ctx context.Context,
) (
	err error,
) {
	err = receiver.ToGateway.ViaGRPC(
		ctx,
	)
	return
}
