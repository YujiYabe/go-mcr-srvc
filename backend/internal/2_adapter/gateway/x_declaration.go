package gateway

import (
	"backend/internal/4_domain/struct_object"
	"context"
)

type (
	Gateway struct {
		ToRedis    ToRedis
		ToPostgres ToPostgres
	}

	// ToRedis ...
	ToRedis interface {
		ResetPlaceListInRedis(
			ctx context.Context,
		) error
	}

	// ToPostgres ...
	ToPostgres interface {
		GetPersonList(
			ctx context.Context,
		) (
			personList struct_object.PersonList,
			err error,
		)

		GetPersonByCondition(
			ctx context.Context,
			reqPerson struct_object.Person,
		) (
			resPersonList struct_object.PersonList,
			err error,
		)
	}
)

// NewGateway ...
func NewGateway(
	toRedis ToRedis,
	toPostgres ToPostgres,
) *Gateway {
	return &Gateway{
		ToRedis:    toRedis,
		ToPostgres: toPostgres,
	}
}
