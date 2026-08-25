package usecase

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
	typeObject "backend/internal/4_domain/type_object"
)

type (

	// ToUseCase ...
	ToUseCase interface {
		Start()

		GetUserList(
			ctx context.Context,
		) (
			userList groupObject.UserList,
			err error,
		)

		GetUserListByCondition(
			ctx context.Context,
			reqUser groupObject.User,
		) (
			resUserList groupObject.UserList,
			err error,
		)

		FetchAccessToken(
			ctx context.Context,
			credential groupObject.Credential,
		) (
			accessToken typeObject.AccessToken,
			err error,
		)

		ViaGRPC(
			ctx context.Context,
			reqUser groupObject.User,
		) (
			resUserList groupObject.UserList,
			err error,
		)

		UpdateUser(
			ctx context.Context,
			newUser groupObject.User,
		) error

		UpdateUserWithEmployment(
			ctx context.Context,
			newUser groupObject.User,
			userEmployment groupObject.UserEmployment,
		) error

		PublishTestTopic(
			ctx context.Context,
		) error
	}
)
