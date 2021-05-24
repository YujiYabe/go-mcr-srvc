package domain

import "context"

type (
	domain struct{}
)

// NewDomain ...
func NewDomain() *domain {
	return &domain{}
}

// Dummy ...
func (domain *domain) Dummy(ctx context.Context) error {
	return nil
}
