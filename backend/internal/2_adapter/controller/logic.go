package controller

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
	typeObject "backend/internal/4_domain/type_object"
)

func (receiver *controller) Start() {}

func (receiver *controller) GetPersonList(
	ctx context.Context,
) (
	personList groupObject.PersonList,
	err error,
) {
	personList, err = receiver.UseCase.GetPersonList(ctx)
	return
}

func (receiver *controller) GetPersonListByCondition(
	ctx context.Context,
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
	err error,
) {
	resPersonList, err = receiver.UseCase.GetPersonListByCondition(
		ctx,
		reqPerson,
	)
	return
}

func (receiver *controller) FetchAccessToken(
	ctx context.Context,
	credential groupObject.Credential,
) (
	accessToken typeObject.AccessToken,
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
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
	err error,
) {
	resPersonList, err = receiver.UseCase.ViaGRPC(
		ctx,
		reqPerson,
	)
	return
}

func (receiver *controller) PublishTestTopic(
	ctx context.Context,
) error {
	return receiver.UseCase.PublishTestTopic(ctx)
}
