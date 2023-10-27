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
	// isDemo := true
	isDemo := false

	return &domain{
		OrderList:      NewOrderList(isDemo),
		Stock:          NewStock(),
		Language:       NewLanguage(),
		AllergyList:    NewAllergyList(),
		AllergyDefault: Allergy(NeAllergyDefault()),
	}
}

// ToDomain ...
type ToDomain interface {
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

	VerifyJANCodes(
		janCodeList []int,
		isVaildJANCodeList []int,
		isVaildLangCodeList []int,
		defaultLangCode int,
	) (
		[]int,
		int,
	)

	GetIsVaildJANCodeList() []int

	// order -----------------------
	GetOrderList(
		ctx context.Context,
	) OrderList

	GetReservingList(
		ctx context.Context,
	) ReservingList

	GetReserving(
		ctx context.Context,
		number int,
	) Reserving

	GetSoldList(
		ctx context.Context,
	) SoldList

	FindSoldIndex(soldNo int) int

	DeleteSoldList(index int)

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

	MergeWithExistingOrder(
		newSold Sold,
	) bool

	AddNewSold(newSold Sold)

	SortOrderList()

	UpdateExistingReserving(
		newReserving Reserving,
	) bool

	AddNewReserving(newReserving Reserving)

	// allergy -----------------------
	GetAllergyDefault(
		ctx context.Context,
	) Allergy

	GetAllergyList(
		ctx context.Context,
	) AllergyList

	// language -----------------------
	GetIsVaildLangCodeMap() map[int]string

	GetIsVaildLangCodeList() []int

	GetDefaultLangCode() int
}
