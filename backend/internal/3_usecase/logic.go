package usecase

import (
	"context"

	structObject "backend/internal/4_domain/struct_object"
	valueObject "backend/internal/4_domain/value_object"
)

// Start ...
func (receiver *useCase) Start() {
}

func (receiver *useCase) GetPersonList(
	ctx context.Context,
) (
	personList structObject.PersonList,
	err error,
) {
	personList, err = receiver.ToGateway.GetPersonList(ctx)
	return
}

func (receiver *useCase) GetPersonByCondition(
	ctx context.Context,
	reqPerson structObject.Person,
) (
	resPersonList structObject.PersonList,
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
	credential structObject.Credential,
) (
	accessToken valueObject.AccessToken,
) {
	accessToken = receiver.ToGateway.FetchAccessToken(
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
