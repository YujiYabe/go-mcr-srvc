package domain

import "context"

type domain struct {
	OrderList
	Stock
	Language
	AllergyList
	AllergyDefault Allergy
}

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

type (
	// ToDomain ...
	ToDomain interface {
		// product -----------------------
		GetProduct(
			ctx context.Context,
			productNumber int,
		) Product

		GetAllProductList(
			ctx context.Context,
		) AllProductList

		GetProductList(
			ctx context.Context,
		) ProductList

		UpdateProduct(
			ctx context.Context,
			product Product,
		)

		SaveInMemory(
			ctx context.Context,
			allProductList AllProductList,
		) error

		// order -----------------------
		GetOrderList(
			ctx context.Context,
		) OrderList

		GetReservingList(
			ctx context.Context,
		) ReservingList

		GetSoldList(
			ctx context.Context,
		) SoldList

		GetPreparingList(
			ctx context.Context,
		) SoldList

		GetCompletedList(
			ctx context.Context,
		) SoldList

		GetPassedList(
			ctx context.Context,
		) SoldList

		UpdateSoldStatus(
			ctx context.Context,
			newSold Sold,
		)

		// allergy -----------------------
		GetAllergyDefault(
			ctx context.Context,
		) Allergy

		GetAllergyList(
			ctx context.Context,
		) AllergyList

		// language -----------------------
		GetIsVaildLangCodeMap(
			ctx context.Context,
		) map[int]string
	}
)
