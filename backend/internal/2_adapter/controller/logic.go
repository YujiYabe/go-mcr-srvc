package controller

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
	typeObject "backend/internal/4_domain/type_object"
)

func (receiver *controller) Start() {}

func (receiver *controller) GetUserList(
	ctx context.Context,
) (
	userList groupObject.UserList,
	err error,
) {
	userList, err = receiver.UseCase.GetUserList(ctx)
	return
}

func (receiver *controller) GetUserListByCondition(
	ctx context.Context,
	reqUser groupObject.User,
) (
	resUserList groupObject.UserList,
	err error,
) {
	resUserList, err = receiver.UseCase.GetUserListByCondition(
		ctx,
		reqUser,
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
	reqUser groupObject.User,
) (
	resUserList groupObject.UserList,
	err error,
) {
	resUserList, err = receiver.UseCase.ViaGRPC(
		ctx,
		reqUser,
	)
	return
}

func (receiver *controller) UpdateUser(
	ctx context.Context,
	newUser groupObject.User,
) error {
	return receiver.UseCase.UpdateUser(ctx, newUser)
}

func (receiver *controller) PublishTestTopic(
	ctx context.Context,
) error {
	return receiver.UseCase.PublishTestTopic(ctx)
}
