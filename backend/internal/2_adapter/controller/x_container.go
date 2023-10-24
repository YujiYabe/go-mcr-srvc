package controller

import (
	"context"

	"backend/internal/2_adapter/gateway"
	"backend/internal/2_adapter/presenter"
	usecase "backend/internal/3_usecase"
	domain "backend/internal/4_domain"

	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("interface_adapter", "controller")
}

type (
	// controller ...
	controller struct {
		UseCase     usecase.ToUseCase
		OrderNumber int
	}

	// orderChannel ...
	orderChannel struct {
		order *domain.Order
	}

	// ToController ...
	ToController interface {
		Start()
		Reserve(ctx context.Context, order *domain.Order, orderType string)
		Order(ctx *context.Context, order *domain.Order)
		GetProduct(
			ctx context.Context,
			productNumber int,
		) *domain.Product
	}
)

// orderChannel ...
var odrChnnl = make(chan orderChannel)

// NewController ...
func NewController(
	toSqlite gateway.ToSqlite,
	// toPostgres gateway.ToPostgres,
	// toMysql gateway.ToMysql,
	// toMongo gateway.ToMongo,
	toShipment presenter.ToShipment,
	toMonitor presenter.ToMonitor,
) ToController {
	toDomain := domain.NewDomain()
	toGateway := gateway.NewGateway(
		toSqlite,
		// toPostgres,
		// toMysql,
		// toMongo,
	)
	toPresenter := presenter.NewPresenter(
		toShipment,
		toMonitor,
	)
	uscs := usecase.NewUseCase(
		toDomain,
		toGateway,
		toPresenter,
	)
	ct := &controller{
		UseCase: uscs,
	}

	return ct
}
