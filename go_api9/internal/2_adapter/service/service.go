package service

// Service ...
type Service struct {
	// ToGrpcOut ToGrpcOut
	ToStocker ToStocker
}

// Dummy ...
func (sv *Service) Dummy() error {
	return nil
}
