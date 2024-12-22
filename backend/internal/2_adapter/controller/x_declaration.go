package controller

import (
	"context"

	"backend/internal/2_adapter/gateway"
	"backend/internal/2_adapter/presenter"
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
	}
)

// NewController ...
func NewController(
	ToRedis gateway.ToRedis,
	ToPostgres gateway.ToPostgres,
	toMySQL gateway.ToMySQL,
	toMongo gateway.ToMongo,
	toShipment presenter.ToShipment,
	toMonitor presenter.ToMonitor,
) (
	toController ToController,
) {
	toDomain := domain.NewDomain()
	toGateway := gateway.NewGateway(
		ToRedis,
		ToPostgres,
		toMySQL,
		toMongo,
	)
	toPresenter := presenter.NewPresenter(
		toShipment,
		toMonitor,
	)
	useCase := usecase.NewUseCase(
		toDomain,
		toGateway,
		toPresenter,
	)

	toController = &controller{
		UseCase: useCase,
	}

	return
}
