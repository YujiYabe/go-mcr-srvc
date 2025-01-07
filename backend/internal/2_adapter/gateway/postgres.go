package gateway

import (
	groupObject "backend/internal/4_domain/group_object"
	"context"
)

// GetPersonList ...
func (receiver *Gateway) GetPersonList(
	ctx context.Context,

) (
	personList groupObject.PersonList,
) {
	return receiver.ToPostgres.GetPersonList(ctx)
}

// GetPersonByCondition ...
func (receiver *Gateway) GetPersonByCondition(
	ctx context.Context,
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
) {
	resPersonList = receiver.ToPostgres.GetPersonByCondition(
		ctx,
		reqPerson,
	)

	return
}
