package usecase

import (
	"backend/internal/4_enterprise_business_rules/entities"
	"context"
)

type (
	// UseCase ...
	UseCase struct {
		ToDomain  ToDomain
		ToService ToService
	}

	// OrderUsecase ...
	OrderUsecase struct {
		ctx   *context.Context
		order *entities.Order
	}

	// ToService ...
	ToService interface {
		GetBans(ctx context.Context, requestBans map[string]int) error
		GetPatties(ctx context.Context, requestPatties map[string]int) error
		GetVegetables(ctx context.Context, requestVegetables map[string]int) error
		GetIngredients(ctx context.Context, requestIngredients map[string]int) error
		Shipment(ctx context.Context, order *entities.Order) error
		UpdateOrders(ctx context.Context, orderNumber string, phase string)
	}

	// ToDomain ...
	ToDomain interface {
		ParseOrder(ctx context.Context, order *entities.Order) *entities.Assemble
		CookHamburgers(ctx context.Context, hamburgers []entities.Hamburger) error
	}
)
