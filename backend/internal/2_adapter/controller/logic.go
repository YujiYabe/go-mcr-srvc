package controller

import (
	"context"

	"backend/internal/4_domain/struct_object"
)

func (receiver *controller) Start() {}

func (receiver *controller) GetPersonList(
	ctx context.Context,
) (
	personList struct_object.PersonList,
	err error,
) {
	return
}
