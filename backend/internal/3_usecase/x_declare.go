package usecase

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
	typeObject "backend/internal/4_domain/type_object"
)

// NewUseCase ...
func NewUseCase(
	toDomain ToDomain,
	toGatewayDB ToGatewayDB,
	toGatewayExternal ToGatewayExternal,
) ToUseCase {
	return &useCase{
		ToDomain:          toDomain,
		ToGatewayDB:       toGatewayDB,
		ToGatewayExternal: toGatewayExternal,
	}
}

type (
	// useCase ...
	useCase struct {
		ToDomain          ToDomain
		ToGatewayDB       ToGatewayDB
		ToGatewayExternal ToGatewayExternal
	}
)

type (
	// ToDomain ...
	ToDomain interface {
		EnsurePrimaryEmploymentAssignable(
			user groupObject.User,
			userEmployment groupObject.UserEmployment,
		) error
	}
)

type (
	// ToGatewayDB ...
	ToGatewayDB interface {
		RunInTransaction(
			ctx context.Context,
			fn func(context.Context) error,
		) error

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

		UpdateUser(
			ctx context.Context,
			newUser groupObject.User,
		) error

		UpdateUserEmployment(
			ctx context.Context,
			userEmployment groupObject.UserEmployment,
		) error

		GetValidationWords(
			ctx context.Context,
			targetType string,
			isBlacklist bool,
		) (
			words []string,
			err error,
		)

		AddValidationWord(
			ctx context.Context,
			targetType string,
			isBlacklist bool,
			word string,
		) error

		UpdateValidationWord(
			ctx context.Context,
			targetType string,
			isBlacklist bool,
			oldWord string,
			newWord string,
		) error

		DeleteValidationWord(
			ctx context.Context,
			targetType string,
			isBlacklist bool,
			word string,
		) error
	}

	// ToGatewayExternal ...
	ToGatewayExternal interface {
		FetchAccessToken(
			ctx context.Context,
			credential groupObject.Credential,
		) (
			accessToken typeObject.AccessToken,
			err error,
		)

		GetUserViaGRPC(
			ctx context.Context,
			reqUser groupObject.User,
		) (
			resUserList groupObject.UserList,
			err error,
		)

		PublishTestTopic(
			ctx context.Context,
		) error
	}
)

type (

	// ToUseCase ...
	ToUseCase interface {
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

		GetUserListViaGRPC(
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

		GetValidationWords(
			ctx context.Context,
			targetType string,
			isBlacklist bool,
		) (
			words []string,
			err error,
		)

		AddValidationWord(
			ctx context.Context,
			targetType string,
			isBlacklist bool,
			word string,
		) error

		UpdateValidationWord(
			ctx context.Context,
			targetType string,
			isBlacklist bool,
			oldWord string,
			newWord string,
		) error

		DeleteValidationWord(
			ctx context.Context,
			targetType string,
			isBlacklist bool,
			word string,
		) error
	}
)
