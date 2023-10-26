package usecase

import (
	"context"

	domain "backend/internal/4_domain"
)

type (
	// useCase ...
	useCase struct {
		ToDomain    domain.ToDomain
		ToGateway   ToGateway
		ToPresenter ToPresenter
	}

	// ToUseCase ...
	ToUseCase interface {
		Start(ctx context.Context)
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

		GetSoldList(
			ctx context.Context,
		) domain.SoldList

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

	// ToGateway ...
	ToGateway interface {
		GetAllProductList(
			ctx context.Context,
		) domain.AllProductList

		UpdateProduct(
			ctx context.Context,
			product domain.Product,
		)
	}

	// ToPresenter ...
	ToPresenter interface {
		UpdateOrders(ctx context.Context)
		Shipment(ctx context.Context) error
		DistributeOrder(
			ctx context.Context,
			orderList *domain.OrderList,
		)
	}
)

// NewUseCase ...
func NewUseCase(
	toDomain domain.ToDomain,
	toGateway ToGateway,
	toPresenter ToPresenter,
) ToUseCase {
	return &useCase{
		ToDomain:    toDomain,
		ToGateway:   toGateway,
		ToPresenter: toPresenter,
	}
}
