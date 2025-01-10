package gateway

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
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
			personList groupObject.PersonList,
		)

		GetPersonListByCondition(
			ctx context.Context,
			reqPerson groupObject.Person,
		) (
			resPersonList groupObject.PersonList,
		)
	}

	// ToAuth0 ...
	ToAuth0 interface {
		FetchAccessToken(
			ctx context.Context,
			credential groupObject.Credential,
		) (
			accessToken valueObject.AccessToken,
		)
	}

	// ToGRPC ...
	ToGRPC interface {
		ViaGRPC(
			ctx context.Context,
			reqPerson groupObject.Person,
		) (
			resPersonList groupObject.PersonList,
		)
	}
)
