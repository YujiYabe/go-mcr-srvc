package db_gateway

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
)

type DBGateway struct {
	ToPostgres ToPostgres
	ToRedis    ToRedis
}

// NewDBGateway ...
func NewDBGateway(
	toPostgres ToPostgres,
	toRedis ToRedis,
) *DBGateway {
	return &DBGateway{
		ToPostgres: toPostgres,
		ToRedis:    toRedis,
	}
}

type (

	// ToPostgres ...
	ToPostgres interface {
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

	// ToRedis ...
	ToRedis interface {
		ResetPlaceListInRedis(
			ctx context.Context,
		) error
	}
)
