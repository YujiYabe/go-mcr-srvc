package external_gateway

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
	valueObject "backend/internal/4_domain/value_object"
)

type GatewayExternal struct {
	ToAuth0 ToAuth0
	ToGRPC  ToGRPC
}

// NewGatewayExternal ...
func NewGatewayExternal(
	toAuth0 ToAuth0,
	toGRPC ToGRPC,
) *GatewayExternal {
	return &GatewayExternal{
		ToAuth0: toAuth0,
		ToGRPC:  toGRPC,
	}
}

type (

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
