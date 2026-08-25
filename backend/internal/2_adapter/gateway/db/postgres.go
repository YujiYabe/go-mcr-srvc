package db_gateway

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
)

func (receiver *GatewayDB) RunInTransaction(
	ctx context.Context,
	fn func(context.Context) error,
) error {
	return receiver.ToPostgres.RunInTransaction(ctx, fn)
}

// GetUserList ...
func (receiver *GatewayDB) GetUserList(
	ctx context.Context,
) (
	userList groupObject.UserList,
	err error,
) {
	return receiver.ToPostgres.GetUserList(
		ctx,
	)
}

// GetUserListByCondition ...
func (receiver *GatewayDB) GetUserListByCondition(
	ctx context.Context,
	reqUser groupObject.User,
) (
	resUserList groupObject.UserList,
	err error,
) {
	resUserList, err = receiver.ToPostgres.GetUserListByCondition(
		ctx,
		reqUser,
	)

	return
}

// UpdateUser ...
func (receiver *GatewayDB) UpdateUser(
	ctx context.Context,
	newUser groupObject.User,
) (
	err error,
) {
	return receiver.ToPostgres.UpdateUser(ctx, newUser)
}
