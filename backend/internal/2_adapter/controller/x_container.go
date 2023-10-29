package controller

import (
	"context"
	"mime/multipart"

	"backend/internal/2_adapter/gateway"
	"backend/internal/2_adapter/presenter"
	usecase "backend/internal/3_usecase"
	domain "backend/internal/4_domain"
)

// controller ...
type controller struct {
	ToUseCase usecase.ToUseCase
}

// NewController ...
func NewController(
	toSqlite gateway.ToSqlite,
	toMonitor presenter.ToMonitor,
	isDemo bool,
) ToController {
	toDomain := domain.NewDomain(isDemo)

	toGateway := gateway.NewGateway(
		toSqlite,
	)

	toPresenter := presenter.NewPresenter(
		toMonitor,
	)

	toUseCase := usecase.NewUseCase(
		toDomain,
		toGateway,
		toPresenter,
	)

	return &controller{
		ToUseCase: toUseCase,
	}
}

// ToController ...
type ToController interface {
	Start()
	// product -----------------------
	GetProduct(
		ctx context.Context,
		productNumber int,
	) domain.Product

	GetAllProductList(
		ctx context.Context,
	) domain.AllProductList

	GetProductList(
		ctx context.Context,
	) domain.ProductList

	UpdateProduct(
		ctx context.Context,
		product domain.Product,
	)

	// order -----------------------
	GetOrderList(
		ctx context.Context,
	) domain.OrderList

	GetReservingList(
		ctx context.Context,
	) domain.ReservingList

	GetReserving(
		ctx context.Context,
		number int,
	) domain.Reserving

	GetSoldList(
		ctx context.Context,
	) domain.SoldList

	SaveSold(
		ctx context.Context,
		newSold domain.Sold,
	)

	SaveReserving(
		ctx context.Context,
		newSold domain.Reserving,
	)

	DeleteSold(
		ctx context.Context,
		number int,
	)

	GetPreparingList(
		ctx context.Context,
	) domain.SoldList

	GetCompletedList(
		ctx context.Context,
	) domain.SoldList

	GetPassedList(
		ctx context.Context,
	) domain.SoldList

	UpdateSoldStatus(
		ctx context.Context,
		newSold domain.Sold,
	)

	DetectSaveJANCodes(
		ctx context.Context,
		number int,
		file *multipart.FileHeader,
	) error

	// allergy -----------------------
	GetAllergyDefault(
		ctx context.Context,
	) domain.Allergy

	GetAllergyList(
		ctx context.Context,
	) domain.AllergyList

	// language -----------------------
	GetIsVaildLangCodeMap(
		ctx context.Context,
	) map[int]string

	// websocket -----------------------
	DistributeOrder(
		ctx context.Context,
	)
}
