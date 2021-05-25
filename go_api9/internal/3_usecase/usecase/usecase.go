package usecase

import "context"

// Order ...
func (uc *UseCase) Order(ctx context.Context) (string, error) {
	res, _ := uc.ToService.Order(ctx)

	return res, nil
}

// Dummy ...
func (uc *UseCase) Dummy(ctx context.Context) (string, error) {
	res, _ := uc.ToService.Dummy(ctx)

	return res, nil
}
