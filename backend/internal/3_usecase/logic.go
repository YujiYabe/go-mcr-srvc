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
	return receiver.ToGateway.GetPersonList(ctx)
}

func (receiver *useCase) GetPersonByCondition(
	ctx context.Context,
	reqPerson struct_object.Person,
) (
	resPersonList struct_object.PersonList,
	err error,
) {
	return receiver.ToGateway.GetPersonByCondition(
		ctx,
		reqPerson,
	)
}

func (receiver *useCase) GetAccessToken(
	ctx context.Context,
	credential struct_object.Credential,
) (
	accessToken value_object.AccessToken,
	err error,
) {
	accessToken, err = receiver.ToGateway.GetAccessToken(
		ctx,
		credential,
	)
	return
}
