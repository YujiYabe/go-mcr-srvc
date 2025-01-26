package usecase

import (
	"context"
	"time"

	groupObject "backend/internal/4_domain/group_object"
	valueObject "backend/internal/4_domain/value_object"
)

// Start ...
func (receiver *useCase) Start() {
}

func (receiver *useCase) GetPersonList(
	ctx context.Context,
) (
	personList groupObject.PersonList,
) {
	personList = receiver.ToDBGateway.GetPersonList(ctx)
	return
}

func (receiver *useCase) GetPersonListByCondition(
	ctx context.Context,
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
) {

	resPersonList = receiver.ToDBGateway.GetPersonListByCondition(
		ctx,
		reqPerson,
	)
	return
}

func (receiver *useCase) FetchAccessToken(
	ctx context.Context,
	credential groupObject.Credential,
) (
	accessToken valueObject.AccessToken,
) {
	accessToken = receiver.ToExternalGateway.FetchAccessToken(
		ctx,
		credential,
	)
	return
}

func (receiver *useCase) ViaGRPC(
	ctx context.Context,
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
) {
	resPersonList = receiver.ToExternalGateway.ViaGRPC(
		ctx,
		reqPerson,
	)

	time.Sleep(20 * time.Second)
	return
}
