package external_gateway

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
)

// ViaGRPC ...
func (receiver *ExternalGateway) ViaGRPC(
	ctx context.Context,
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
) {
	resPersonList = receiver.ToGRPC.ViaGRPC(
		ctx,
		reqPerson,
	)
	return
}
