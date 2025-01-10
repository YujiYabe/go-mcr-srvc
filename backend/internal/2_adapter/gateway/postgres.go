package gateway

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
)

// GetPersonList ...
func (receiver *Gateway) GetPersonList(
	ctx context.Context,

) (
	personList groupObject.PersonList,
) {
	return receiver.ToPostgres.GetPersonList(ctx)
}

// GetPersonListByCondition ...
func (receiver *Gateway) GetPersonListByCondition(
	ctx context.Context,
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
) {
	resPersonList = receiver.ToPostgres.GetPersonListByCondition(
		ctx,
		reqPerson,
	)

	return
}
