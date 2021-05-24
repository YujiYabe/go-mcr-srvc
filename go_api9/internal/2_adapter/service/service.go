package service

import "context"

// Service ...
type Service struct {
	// ToGrpcOut ToGrpcOut
	ToStocker ToStocker
}

// Dummy ...
func (sv *Service) Dummy(ctx context.Context) error {
	sv.ToStocker.Dummy(ctx)
	return nil
}
