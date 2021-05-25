package service

import "context"

// Service ...
type Service struct {
	// ToGrpcOut ToGrpcOut
	ToStocker ToStocker
}

// Order ...
func (sv *Service) Order(ctx context.Context) (string, error) {
	res, _ := sv.ToStocker.Dummy(ctx)

	return res, nil
}

// Dummy ...
func (sv *Service) Dummy(ctx context.Context) (string, error) {
	res, _ := sv.ToStocker.Dummy(ctx)

	return res, nil
}
