package usecase

import (
	"backend/internal/4_enterprise_business_rule/entity"
	"context"
)

type (
	// UseCase ...
	UseCase struct {
		ToEntity    ToEntity
		ToGateway   ToGateway
		ToPresenter ToPresenter
	}

	// ToUseCase ...
	ToUseCase interface {
		Reserve(ctx context.Context, orderinfo *entity.OrderInfo)
		Order(ctx *context.Context, order *entity.Order) error
	}

	// OrderUsecase ...
	OrderUsecase struct {
		ctx   *context.Context
		order *entity.Order
	}

	// ToEntity ...
	ToEntity interface {
		ParseOrder(ctx context.Context, order *entity.Order) *entity.Assemble
		CookHamburgers(ctx context.Context, hamburgers []entity.Hamburger) error
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
		Shipment(ctx context.Context, order *entity.Order) error
	}
)
