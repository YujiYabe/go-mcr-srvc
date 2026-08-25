package usecase

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
	typeObject "backend/internal/4_domain/type_object"
)

// Start ...
func (receiver *useCase) Start() {
}

func (receiver *useCase) GetPersonList(
	ctx context.Context,
) (
	personList groupObject.PersonList,
	err error,
) {
	personList, err = receiver.ToGatewayDB.GetPersonList(ctx)
	return
}

func (receiver *useCase) GetPersonListByCondition(
	ctx context.Context,
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
	err error,
) {

	resPersonList, err = receiver.ToGatewayDB.GetPersonListByCondition(
		ctx,
		reqPerson,
	)
	return
}

func (receiver *useCase) FetchAccessToken(
	ctx context.Context,
	credential groupObject.Credential,
) (
	accessToken typeObject.AccessToken,
	err error,
) {
	accessToken, err = receiver.ToGatewayExternal.FetchAccessToken(
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
	err error,
) {
	resPersonList, err = receiver.ToGatewayExternal.ViaGRPC(
		ctx,
		reqPerson,
	)
	return
}

func (receiver *useCase) PublishTestTopic(
	ctx context.Context,
) {
	receiver.ToGatewayExternal.PublishTestTopic(ctx)
}
