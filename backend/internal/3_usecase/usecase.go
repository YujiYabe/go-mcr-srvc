package usecase

import (
	"context"
	"fmt"

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

		UpdateUserProfileWithPrimaryEmployment(
			ctx context.Context,
			newUser groupObject.User,
			userEmployment groupObject.UserEmployment,
		) error

		PublishTestTopic(
			ctx context.Context,
		) error
	}
)

// Start ...
func (receiver *useCase) Start() {
}

func ensureContextReady(
	ctx context.Context,
	usecaseName string,
) error {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("%s: context is not ready: %w", usecaseName, err)
	}

	return nil
}
