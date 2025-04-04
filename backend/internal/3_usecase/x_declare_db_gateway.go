package usecase

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
)

type (

	// ToGatewayDB ...
	ToGatewayDB interface {
		GetPersonList(
			ctx context.Context,
		) (
			personList groupObject.PersonList,
		)

		GetPersonListByCondition(
			ctx context.Context,
			reqPerson groupObject.Person,
		) (
			resPersonList groupObject.PersonList,
		)
	}
)
