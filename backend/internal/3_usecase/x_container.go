package usecase

import (
	"context"

	domain "backend/internal/4_domain"
)

type (
	// useCase ...
	useCase struct {
		ToDomain    ToDomain
		ToGateway   ToGateway
		ToPresenter ToPresenter
	}

	// ToUseCase ...
	ToUseCase interface {
		Start()
		Reserve(ctx context.Context)
		Order(ctx *context.Context) error
		GetProduct(
			ctx context.Context,
			productNumber int,
		) *domain.Product
	}

	// OrderUseCase ...
	OrderUseCase struct {
		ctx   *context.Context
		order *domain.Order
	}

	// ToDomain ...
	ToDomain interface {
		ParseOrder(ctx context.Context) error
		CookHamburgers(ctx context.Context) error
		GetProduct(
			ctx context.Context,
			productNumber int,
		) *domain.Product
	}

	// ToGateway ...
	ToGateway interface {
		GetBans(ctx context.Context) error
		GetPatties(ctx context.Context) error
		GetVegetables(ctx context.Context) error
		GetIngredients(ctx context.Context) error
		GetAllProductList(
			ctx context.Context,
		) *domain.AllProductList
	}

	// ToPresenter ...
	ToPresenter interface {
		UpdateOrders(ctx context.Context)
		Shipment(ctx context.Context) error
	}
)

// NewUseCase ...
func NewUseCase(
	toDomain ToDomain,
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
