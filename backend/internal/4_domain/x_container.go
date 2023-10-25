package domain

import "context"

type (
	domain struct {
		OrderList
		Stock
		Language
		AllergyList
		AllergyDefault Allergy
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
		) Allergy

		GetSoldList(
			ctx context.Context,
		) SoldList

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
	isDemo := true
	// isDemo := false

	return &domain{
		OrderList:      NewOrderList(isDemo),
		Stock:          NewStock(),
		Language:       NewLanguage(),
		AllergyList:    NewAllergyList(),
		AllergyDefault: Allergy(NeAllergyDefault()),
	}
}
