package domain

type (
	domain struct{}

	// ToDomain ...
	ToDomain interface {
	}
)

// NewDomain ...
func NewDomain() ToDomain {
	return &domain{}
}
