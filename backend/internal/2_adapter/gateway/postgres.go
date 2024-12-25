package gateway

import (
	"backend/internal/4_domain/struct_object"
	"context"
)

// GetPersonList ...
func (receiver *Gateway) GetPersonList(
	ctx context.Context,

) (
	personList struct_object.PersonList,
	err error,
) {
	return receiver.ToPostgres.GetPersonList(ctx)
}

// GetPersonByCondition ...
func (receiver *Gateway) GetPersonByCondition(
	ctx context.Context,
	reqPerson struct_object.Person,
) (
	resPersonList struct_object.PersonList,
	err error,
) {
	return receiver.ToPostgres.GetPersonByCondition(
		ctx,
		reqPerson,
	)
}
