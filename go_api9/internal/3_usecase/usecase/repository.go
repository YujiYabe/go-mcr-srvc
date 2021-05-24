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
		Dummy(ctx context.Context) error
	}

	// ToDomain ...
	ToDomain interface {
		Dummy(ctx context.Context) error
	}
)
