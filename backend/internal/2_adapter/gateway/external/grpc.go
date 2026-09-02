package external_gateway

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
)

// GetUserViaGRPC ...
func (receiver *GatewayExternal) GetUserViaGRPC(
	ctx context.Context,
	reqUser groupObject.User,
) (
	resUserList groupObject.UserList,
	err error,
) {
	resUserList, err = receiver.ToGRPC.GetUserViaGRPC(
		ctx,
		reqUser,
	)

	return
}
