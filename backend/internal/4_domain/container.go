package domain

import (
	"context"
)

type (
	domain struct{}

	// ToDomain ...
	ToDomain interface {
		ParseOrder(ctx context.Context) error
		CookHamburgers(ctx context.Context) error
	}
)

type (
	Order struct {
	}
)
