package controller

import (
	"context"

	externalGateway "backend/internal/2_adapter/gateway/external"
	// externalGateway "backend/internal/2_adapter/gateway/external"
	dbGateway "backend/internal/2_adapter/gateway/db"
	usecase "backend/internal/3_usecase"

	domain "backend/internal/4_domain"
	groupObject "backend/internal/4_domain/group_object"
	valueObject "backend/internal/4_domain/value_object"
)

// NewController ...
func NewController(
	ToRedis dbGateway.ToRedis,
	ToPostgres dbGateway.ToPostgres,
	ToAuth0 externalGateway.ToAuth0,
	ToGRPC externalGateway.ToGRPC,
) (
	toController ToController,
) {
	toDomain := domain.NewDomain()

	toDBGateway := dbGateway.NewDBGateway(
		ToPostgres,
		ToRedis,
	)

	toExternalGateway := externalGateway.NewExternalGateway(
		ToAuth0,
		ToGRPC,
	)

	useCase := usecase.NewUseCase(
		toDomain,
		toDBGateway,
		toExternalGateway,
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
	}
)
