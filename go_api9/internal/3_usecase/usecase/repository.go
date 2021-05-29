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
		GetVegetables(ctx context.Context, requestVegetables map[string]int) error
		Dummy(ctx context.Context) (string, error)
	}

	// ToDomain ...
	ToDomain interface {
		ParseOrder(ctx context.Context, order domain.Order) *domain.Assemble
		Dummy(ctx context.Context) error
	}
)
