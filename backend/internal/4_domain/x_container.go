package domain

import "context"

type (
	Order struct {
	}

	domain struct {
		*Stock
		*Language
		*AllergyList
		*AllergyDefault
	}

	// ToDomain ...
	ToDomain interface {
		ParseOrder(ctx context.Context) error
		CookHamburgers(ctx context.Context) error
		GetProduct(
			ctx context.Context,
			productNumber int,
		) *Product

		GetAllProductList(
			ctx context.Context,
		) *AllProductList

		GetAllergyDefault(
			ctx context.Context,
		) *Allergy

		GetIsVaildLangCodeMap(
			ctx context.Context,
		) map[int]string

		SaveInMemory(
			ctx context.Context,
			allProductList *AllProductList,
		) error
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
