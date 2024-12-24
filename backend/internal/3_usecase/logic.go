package usecase

import (
	"backend/internal/4_domain/struct_object"
	"context"
)

// Start ...
func (receiver *useCase) Start() {
}

func (receiver *useCase) GetPersonList(
	ctx context.Context,
) (
	personList struct_object.PersonList,
	err error,
) {
	return receiver.ToGateway.GetPersonList(ctx)
}

func (receiver *useCase) GetPersonByCondition(
	ctx context.Context,
	reqPerson struct_object.Person,
) (
	resPersonList struct_object.PersonList,
	err error,
) {
	return receiver.ToGateway.GetPersonByCondition(
		ctx,
		reqPerson,
	)
}
