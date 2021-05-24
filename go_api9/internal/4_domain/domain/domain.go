package domain

type (
	domain struct{}
)

// NewDomain ...
func NewDomain() *domain {
	return &domain{}
}

// Dummy ...
func (domain *domain) Dummy() error {
	return nil
}
