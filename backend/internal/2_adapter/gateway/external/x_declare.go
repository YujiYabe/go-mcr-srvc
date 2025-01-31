package external_gateway

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
	valueObject "backend/internal/4_domain/value_object"
)

type GatewayExternal struct {
	ToAuth0  ToAuth0
	ToGRPC   ToGRPC
	ToPubSub ToPubSub
}

// NewGatewayExternal ...
func NewGatewayExternal(
	toAuth0 ToAuth0,
	toGRPC ToGRPC,
	toPubSub ToPubSub,
) *GatewayExternal {
	return &GatewayExternal{
		ToAuth0:  toAuth0,
		ToGRPC:   toGRPC,
		ToPubSub: toPubSub,
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

	// ToPubSub ...
	ToPubSub interface {
		PublishTestTopic(
			ctx context.Context,
		)
	}
)
