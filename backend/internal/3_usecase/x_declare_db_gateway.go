package usecase

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
)

type (

	// ToGatewayDB ...
	ToGatewayDB interface {
		RunInTransaction(
			ctx context.Context,
			fn func(context.Context) error,
		) error

		GetPersonList(
			ctx context.Context,
		) (
			personList groupObject.PersonList,
			err error,
		)

		GetPersonListByCondition(
			ctx context.Context,
			reqPerson groupObject.Person,
		) (
			resPersonList groupObject.PersonList,
			err error,
		)

		UpdatePerson(
			ctx context.Context,
			newPerson groupObject.Person,
		) error
	}
)
