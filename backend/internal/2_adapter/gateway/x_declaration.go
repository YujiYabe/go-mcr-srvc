package gateway

import (
	"context"

	"backend/internal/4_domain/struct_object"
	"backend/internal/4_domain/value_object"
)

type (
	Gateway struct {
		ToRedis    ToRedis
		ToPostgres ToPostgres
		ToAuth0    ToAuth0
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

	ToAuth0 interface {
		GetAccessToken(
			ctx context.Context,
			credential struct_object.Credential,
		) (
			accessToken value_object.AccessToken,
			err error,
		)
	}
)

// NewGateway ...
func NewGateway(
	toRedis ToRedis,
	toPostgres ToPostgres,
	toAuth0 ToAuth0,
) *Gateway {
	return &Gateway{
		ToRedis:    toRedis,
		ToPostgres: toPostgres,
		ToAuth0:    toAuth0,
	}
}
