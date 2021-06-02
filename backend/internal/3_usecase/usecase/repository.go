package usecase

import (
	"app/internal/4_domain/domain"
	"context"
)

type (
	// UseCase ...
	UseCase struct {
		ToDomain  ToDomain
		ToService ToService
	}

	// ToService ...
	ToService interface {
		GetBans(ctx context.Context, requestBans map[string]int) error
		GetPatties(ctx context.Context, requestPatties map[string]int) error
		GetVegetables(ctx context.Context, requestVegetables map[string]int) error
		GetIngredients(ctx context.Context, requestIngredients map[string]int) error
		Shipment(ctx context.Context, order *domain.Order) error
		Dummy(ctx context.Context) (string, error)
	}

	// ToDomain ...
	ToDomain interface {
		ParseOrder(ctx context.Context, order *domain.Order) *domain.Assemble
		Dummy(ctx context.Context) error
	}
)
