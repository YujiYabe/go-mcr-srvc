package controller

import (
	"context"

	"backend/internal/2_adapter/gateway"
	usecase "backend/internal/3_usecase"

	domain "backend/internal/4_domain"
	groupObject "backend/internal/4_domain/group_object"
	valueObject "backend/internal/4_domain/value_object"
)

// NewController ...
func NewController(
	ToRedis gateway.ToRedis,
	ToPostgres gateway.ToPostgres,
	ToAuth0 gateway.ToAuth0,
	ToGRPC gateway.ToGRPC,
) (
	toController ToController,
) {
	toDomain := domain.NewDomain()
	toGateway := gateway.NewGateway(
		ToRedis,
		ToPostgres,
		ToAuth0,
		ToGRPC,
	)

	useCase := usecase.NewUseCase(
		toDomain,
		toGateway,
	)

	toController = &controller{
		UseCase: useCase,
	}

	return
}

type (
	// controller ...
	controller struct {
		UseCase     usecase.ToUseCase
		OrderNumber int
	}

	// ToController ...
	ToController interface {
		Start()

		GetPersonList(
			ctx context.Context,
		) (
			personList groupObject.PersonList,
		)

		GetPersonByCondition(
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
		) (
			err error,
		)
	}
)
