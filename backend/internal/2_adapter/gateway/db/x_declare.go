package db_gateway

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
	typeObject "backend/internal/4_domain/type_object"
)

type GatewayDB struct {
	ToPostgres ToPostgres
	ToRedis    ToRedis
}

// NewGatewayDB ...
func NewGatewayDB(
	toPostgres ToPostgres,
	toRedis ToRedis,
) *GatewayDB {
	return &GatewayDB{
		ToPostgres: toPostgres,
		ToRedis:    toRedis,
	}
}

type (

	// ToPostgres ...
	ToPostgres interface {
		GetPerson(
			ctx context.Context,
			id typeObject.ID,
		) (
			person groupObject.Person,
			err error,
		)

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
	}

	// ToRedis ...
	ToRedis interface {
		ResetPlaceListInRedis(
			ctx context.Context,
		) error
	}
)
