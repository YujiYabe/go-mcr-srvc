package usecase

import (
	domain "backend/internal/4_domain"
	"context"
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
		Reserve(ctx context.Context, orderInfo *domain.OrderInfo)
		Order(ctx *context.Context, order *domain.Order) error
	}
	//

	// OrderUseCase ...
	OrderUseCase struct {
		ctx   *context.Context
		order *domain.Order
	}

	// ToDomain ...
	ToDomain interface {
		ParseOrder(ctx context.Context, order *domain.Order) *domain.Assemble
		CookHamburgers(ctx context.Context, hamburgers []domain.Hamburger) error
	}

	// ToGateway ...
	ToGateway interface {
		GetBans(ctx context.Context, requestBans map[string]int) error
		GetPatties(ctx context.Context, requestPatties map[string]int) error
		GetVegetables(ctx context.Context, requestVegetables map[string]int) error
		GetIngredients(ctx context.Context, requestIngredients map[string]int) error
	}

	// ToPresenter ...
	ToPresenter interface {
		UpdateOrders(ctx context.Context, orderNumber string, phase string)
		Shipment(ctx context.Context, order *domain.Order) error
	}
)

// NewUseCase ...
func NewUseCase(
	toDomain ToDomain,
	toGateway ToGateway,
	toPresenter ToPresenter,
) ToUseCase {
	return &useCase{
		ToDomain:    toDomain,
		ToGateway:   toGateway,
		ToPresenter: toPresenter,
	}
}
