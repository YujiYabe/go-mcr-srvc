package usecase

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
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
	}
)
