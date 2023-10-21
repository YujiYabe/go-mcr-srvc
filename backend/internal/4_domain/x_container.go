package domain

import (
	"context"
)

type (
	Order struct {
	}

	domain struct {
		*Language
		*AllergyList
		*Stock
	}

	// ToDomain ...
	ToDomain interface {
		ParseOrder(ctx context.Context) error
		CookHamburgers(ctx context.Context) error
	}
)

// NewDomain ...
func NewDomain() ToDomain {
	return &domain{
		Stock:       NewStock(),
		Language:    NewLanguage(),
		AllergyList: NewAllergyList(),
	}
}
