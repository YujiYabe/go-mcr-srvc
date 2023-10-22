package usecase

import (
	"context"

	"backend/pkg"
)

var (
	myErr        *pkg.MyErr
	orderUseCase = make(chan OrderUseCase)
)

func init() {
	myErr = pkg.NewMyErr("application_business_rule", "usecase")
}

// Start ...
func (uscs *useCase) Start() {
	go uscs.bulkOrder()
}

// Reserve ...
func (uscs *useCase) Reserve(ctx context.Context) {}

// Order ...
func (uscs *useCase) Order(ctx *context.Context) error {

	return nil
}

func (uscs *useCase) bulkOrder() {}
