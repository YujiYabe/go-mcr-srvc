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

		GetIsVaildLangCodeMap(
			ctx context.Context,
		) map[int]string
		SaveInMemory(
			ctx context.Context,
			allProductList AllProductList,
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
