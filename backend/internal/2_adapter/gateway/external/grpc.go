package external_gateway

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
)

// ViaGRPC ...
func (receiver *GatewayExternal) ViaGRPC(
	ctx context.Context,
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
	err error,
) {
	resPersonList, err = receiver.ToGRPC.ViaGRPC(
		ctx,
		reqPerson,
	)
	return
}
