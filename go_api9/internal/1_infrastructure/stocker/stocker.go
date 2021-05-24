package stocker

import (
	"app/internal/2_adapter/service"
	"context"
)

type stocker struct{}

// NewToStocker ...
func NewToStocker() service.ToStocker {
	s := new(stocker)
	return s
}

// Dummy ...
func (out *stocker) Dummy(ctx context.Context) error {
	return nil
}
