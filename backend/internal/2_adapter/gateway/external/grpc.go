package external_gateway

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
)

// ViaGRPC ...
func (receiver *GatewayExternal) ViaGRPC(
	ctx context.Context,
	reqUser groupObject.User,
) (
	resUserList groupObject.UserList,
	err error,
) {
	resUserList, err = receiver.ToGRPC.ViaGRPC(
		ctx,
		reqUser,
	)
	return
}
