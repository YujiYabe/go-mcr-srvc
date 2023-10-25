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

		GetIsVaildLangCodeMap(
			ctx context.Context,
		) map[int]string
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
	}
)

// NewUseCase ...
func NewUseCase(
	toDomain domain.ToDomain,
	toGateway ToGateway,
	toPresenter ToPresenter,
) ToUseCase {

	uscs := &useCase{
		ToDomain:    toDomain,
		ToGateway:   toGateway,
		ToPresenter: toPresenter,
	}

	return uscs
}
