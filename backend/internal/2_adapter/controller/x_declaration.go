package controller

import (
	"context"

	"backend/internal/2_adapter/gateway"
	usecase "backend/internal/3_usecase"
	domain "backend/internal/4_domain"
	"backend/internal/4_domain/struct_object"
)

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

// NewController ...
func NewController(
	ToRedis gateway.ToRedis,
	ToPostgres gateway.ToPostgres,
) (
	toController ToController,
) {
	toDomain := domain.NewDomain()
	toGateway := gateway.NewGateway(
		ToRedis,
		ToPostgres,
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
