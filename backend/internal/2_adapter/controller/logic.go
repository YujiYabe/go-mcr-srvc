package controller

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
	domainObject "backend/internal/4_domain/type_object"
)

func (receiver *controller) Start() {}

func (receiver *controller) GetPersonList(
	ctx context.Context,
) (
	personList groupObject.PersonList,
) {
	personList = receiver.UseCase.GetPersonList(ctx)
	return
}

func (receiver *controller) GetPersonListByCondition(
	ctx context.Context,
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
) {
	resPersonList = receiver.UseCase.GetPersonListByCondition(
		ctx,
		reqPerson,
	)
	return
}

func (receiver *controller) FetchAccessToken(
	ctx context.Context,
	credential groupObject.Credential,
) (
	accessToken domainObject.AccessToken,
) {
	accessToken = receiver.UseCase.FetchAccessToken(
		ctx,
		credential,
	)
	return
}

func (receiver *controller) ViaGRPC(
	ctx context.Context,
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
) {
	resPersonList = receiver.UseCase.ViaGRPC(
		ctx,
		reqPerson,
	)
	return
}

func (receiver *controller) PublishTestTopic(
	ctx context.Context,
) {
	receiver.UseCase.PublishTestTopic(ctx)
}
