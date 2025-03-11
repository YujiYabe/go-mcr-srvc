package usecase

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
	domainObject "backend/internal/4_domain/type_object"
)

type (

	// ToGatewayExternal ...
	ToGatewayExternal interface {
		FetchAccessToken(
			ctx context.Context,
			credential groupObject.Credential,
		) (
			accessToken domainObject.AccessToken,
		)

		ViaGRPC(
			ctx context.Context,
			reqPerson groupObject.Person,
		) (
			resPersonList groupObject.PersonList,
		)

		PublishTestTopic(
			ctx context.Context,
		)
	}
)
