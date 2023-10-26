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
		// product -----------------------
		GetProduct(
			ctx context.Context,
			productNumber int,
		) domain.Product

		GetAllProductList(
			ctx context.Context,
		) domain.AllProductList

		GetProductList(
			ctx context.Context,
		) domain.ProductList

		UpdateProduct(
			ctx context.Context,
			product domain.Product,
		)

		// order -----------------------
		GetOrderList(
			ctx context.Context,
		) domain.OrderList

		GetReservingList(
			ctx context.Context,
		) domain.ReservingList

		GetReserving(
			ctx context.Context,
			number int,
		) domain.Reserving

		GetSoldList(
			ctx context.Context,
		) domain.SoldList

		SaveSold(
			ctx context.Context,
			newSold domain.Sold,
		)

		DeleteSold(
			ctx context.Context,
			number int,
		)

		GetPreparingList(
			ctx context.Context,
		) domain.SoldList

		GetCompletedList(
			ctx context.Context,
		) domain.SoldList

		GetPassedList(
			ctx context.Context,
		) domain.SoldList

		UpdateSoldStatus(
			ctx context.Context,
			newSold domain.Sold,
		)

		// allergy -----------------------
		GetAllergyDefault(
			ctx context.Context,
		) domain.Allergy

		GetAllergyList(
			ctx context.Context,
		) domain.AllergyList

		// language -----------------------
		GetIsVaildLangCodeMap(
			ctx context.Context,
		) map[int]string

		// websocket -----------------------
		DistributeOrder(
			ctx context.Context,
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
