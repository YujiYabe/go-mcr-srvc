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
		ToUseCase usecase.ToUseCase
	}

	// ToController ...
	ToController interface {
		Start()

		GetProduct(
			ctx context.Context,
			productNumber int,
		) *domain.Product

		GetAllProductList(
			ctx context.Context,
		) *domain.AllProductList

		GetAllergyDefault(
			ctx context.Context,
		) domain.Allergy

		GetIsVaildLangCodeMap(
			ctx context.Context,
		) map[int]string

		UpdateProduct(
			ctx context.Context,
			product domain.Product,
		)
	}
)

// NewController ...
func NewController(
	toSqlite gateway.ToSqlite,
	toShipment presenter.ToShipment,
	toMonitor presenter.ToMonitor,
) ToController {
	toDomain := domain.NewDomain()

	toGateway := gateway.NewGateway(
		toSqlite,
	)

	toPresenter := presenter.NewPresenter(
		toShipment,
		toMonitor,
	)

	toUseCase := usecase.NewUseCase(
		toDomain,
		toGateway,
		toPresenter,
	)

	return &controller{
		ToUseCase: toUseCase,
	}
}
