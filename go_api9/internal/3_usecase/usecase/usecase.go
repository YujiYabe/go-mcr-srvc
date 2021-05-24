package usecase

import "context"

// Dummy ...
func (uc *UseCase) Dummy(ctx context.Context) error {
	uc.ToService.Dummy(ctx)
	return nil
}
