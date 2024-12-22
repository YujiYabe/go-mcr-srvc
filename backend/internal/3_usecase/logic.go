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

	return
}
