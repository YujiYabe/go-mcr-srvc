package gateway

import (
	"context"

	structObject "backend/internal/4_domain/struct_object"
)

// GetPersonList ...
func (receiver *Gateway) GetPersonList(
	ctx context.Context,

) (
	personList structObject.PersonList,
	err error,
) {
	return receiver.ToPostgres.GetPersonList(ctx)
}

// GetPersonByCondition ...
func (receiver *Gateway) GetPersonByCondition(
	ctx context.Context,
	reqPerson structObject.Person,
) (
	resPersonList structObject.PersonList,
	err error,
) {
	resPersonList, err = receiver.ToPostgres.GetPersonByCondition(
		ctx,
		reqPerson,
	)
	return
}
