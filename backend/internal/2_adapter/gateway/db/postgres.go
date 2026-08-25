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

// GetPersonList ...
func (receiver *GatewayDB) GetPersonList(
	ctx context.Context,
) (
	personList groupObject.PersonList,
	err error,
) {
	return receiver.ToPostgres.GetPersonList(
		ctx,
	)
}

// GetPersonListByCondition ...
func (receiver *GatewayDB) GetPersonListByCondition(
	ctx context.Context,
	reqPerson groupObject.Person,
) (
	resPersonList groupObject.PersonList,
	err error,
) {
	resPersonList, err = receiver.ToPostgres.GetPersonListByCondition(
		ctx,
		reqPerson,
	)

	return
}

// UpdatePerson ...
func (receiver *GatewayDB) UpdatePerson(
	ctx context.Context,
	newPerson groupObject.Person,
) (
	err error,
) {
	_, err = receiver.ToPostgres.GetPerson(
		ctx,
		newPerson.ID(),
	)
	return
}
