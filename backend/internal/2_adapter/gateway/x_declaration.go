package gateway

import (
	"context"

	structObject "backend/internal/4_domain/struct_object"
	valueObject "backend/internal/4_domain/value_object"
)

type Gateway struct {
	ToRedis    ToRedis
	ToPostgres ToPostgres
	ToAuth0    ToAuth0
	ToGRPC     ToGRPC
}

// NewGateway ...
func NewGateway(
	toRedis ToRedis,
	toPostgres ToPostgres,
	toAuth0 ToAuth0,
	toGRPC ToGRPC,
) *Gateway {
	return &Gateway{
		ToRedis:    toRedis,
		ToPostgres: toPostgres,
		ToAuth0:    toAuth0,
		ToGRPC:     toGRPC,
	}
}

type (
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
			personList structObject.PersonList,
			err error,
		)

		GetPersonByCondition(
			ctx context.Context,
			reqPerson structObject.Person,
		) (
			resPersonList structObject.PersonList,
			err error,
		)
	}

	// ToAuth0 ...
	ToAuth0 interface {
		FetchAccessToken(
			ctx context.Context,
			credential structObject.Credential,
		) (
			accessToken valueObject.AccessToken,
			err error,
		)
	}

	// ToGRPC ...
	ToGRPC interface {
		ViaGRPC(
			ctx context.Context,
		) (
			err error,
		)
	}
)
