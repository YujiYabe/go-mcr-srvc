package usecase

import (
	"context"
	"time"

	domainObject "backend/internal/4_domain/domain_object"
	groupObject "backend/internal/4_domain/group_object"
)

// Start ...
func (receiver *useCase) Start() {
}

func (receiver *useCase) GetPersonList(
	ctx context.Context,
) (
	personList groupObject.PersonList,
) {
	personList = receiver.ToGatewayDB.GetPersonList(ctx)
	return
}

func (receiver *useCase) GetPersonListByCondition(
	ctx context.Context,
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
) {

	resPersonList = receiver.ToGatewayDB.GetPersonListByCondition(
		ctx,
		reqPerson,
	)
	return
}

func (receiver *useCase) FetchAccessToken(
	ctx context.Context,
	credential groupObject.Credential,
) (
	accessToken domainObject.AccessToken,
) {
	accessToken = receiver.ToGatewayExternal.FetchAccessToken(
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
	resPersonList = receiver.ToGatewayExternal.ViaGRPC(
		ctx,
		reqPerson,
	)

	time.Sleep(20 * time.Second)
	return
}

func (receiver *useCase) PublishTestTopic(
	ctx context.Context,
) {
	receiver.ToGatewayExternal.PublishTestTopic(ctx)
}
