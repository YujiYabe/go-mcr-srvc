package db_gateway

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
	typeObject "backend/internal/4_domain/type_object"
)

type GatewayDB struct {
	ToPostgres ToPostgres
	ToRedis    ToRedis
}

// NewGatewayDB ...
func NewGatewayDB(
	toPostgres ToPostgres,
	toRedis ToRedis,
) *GatewayDB {
	return &GatewayDB{
		ToPostgres: toPostgres,
		ToRedis:    toRedis,
	}
}

type (

	// ToPostgres ...
	ToPostgres interface {
		RunInTransaction(
			ctx context.Context,
			fn func(context.Context) error,
		) error

		GetUser(
			ctx context.Context,
			id typeObject.ID,
		) (
			user groupObject.User,
			err error,
		)

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

	// ToRedis ...
	ToRedis interface {
		ResetPlaceListInRedis(
			ctx context.Context,
		) error
	}
)
