package controller

import (
	"context"

	gatewayDB "backend/internal/2_adapter/gateway/db"
	gatewayExternal "backend/internal/2_adapter/gateway/external"
	usecase "backend/internal/3_usecase"

	domain "backend/internal/4_domain"
	groupObject "backend/internal/4_domain/group_object"
	valueObject "backend/internal/4_domain/value_object"
)

// NewController ...
func NewController(
	ToPostgres gatewayDB.ToPostgres,
	ToRedis gatewayDB.ToRedis,
	ToAuth0 gatewayExternal.ToAuth0,
	ToGRPC gatewayExternal.ToGRPC,
	ToPubSub gatewayExternal.ToPubSub,
) (
	toController ToController,
) {
	toDomain := domain.NewDomain()

	toGatewayDB := gatewayDB.NewGatewayDB(
		ToPostgres,
		ToRedis,
	)

	toGatewayExternal := gatewayExternal.NewGatewayExternal(
		ToAuth0,
		ToGRPC,
		ToPubSub,
	)

	useCase := usecase.NewUseCase(
		toDomain,
		toGatewayDB,
		toGatewayExternal,
	)

	toController = &controller{
		UseCase: useCase,
	}

	return
}

type (
	// controller ...
	controller struct {
		UseCase usecase.ToUseCase
	}

	// ToController ...
	ToController interface {
		Start()

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

		FetchAccessToken(
			ctx context.Context,
			credential groupObject.Credential,
		) (
			accessToken valueObject.AccessToken,
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
