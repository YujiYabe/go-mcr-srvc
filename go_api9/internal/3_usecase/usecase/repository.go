package usecase

import "context"

type (
	// UseCase ...
	UseCase struct {
		ToDomain  ToDomain
		ToService ToService
	}

	// ToService ...
	ToService interface {
		Dummy(ctx context.Context) (string, error)
		Order(ctx context.Context) (string, error)
		GetVegetables(ctx context.Context, requestVegetables map[string]int) error
	}

	// ToDomain ...
	ToDomain interface {
		Dummy(ctx context.Context) error
	}
)
