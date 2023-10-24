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
		Start(ctx *context.Context)
		Reserve(ctx context.Context)
		Order(ctx *context.Context) error

		GetProduct(
			ctx context.Context,
			productNumber int,
		) *domain.Product

		SetUpInMemory(
			ctx *context.Context,
		)
	}

	// OrderUseCase ...
	OrderUseCase struct {
		ctx   *context.Context
		order *domain.Order
	}

	// ToGateway ...
	ToGateway interface {
		GetBans(ctx context.Context) error
		GetPatties(ctx context.Context) error
		GetVegetables(ctx context.Context) error
		GetIngredients(ctx context.Context) error
		GetAllProductList(
			ctx *context.Context,
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
