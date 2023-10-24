package usecase

import (
	"context"

	domain "backend/internal/4_domain"
)

type (
	// useCase ...
	useCase struct {
		ToDomain    domain.ToDomain
		ToGateway   ToGateway
		ToPresenter ToPresenter
	}

	// ToUseCase ...
	ToUseCase interface {
		Start(ctx context.Context)
		Reserve(ctx context.Context)
		Order(ctx context.Context) error

		GetProduct(
			ctx context.Context,
			productNumber int,
		) *domain.Product

		GetAllergyDefault(
			ctx context.Context,
		) domain.Allergy

		GetAllProductList(
			ctx context.Context,
		) *domain.AllProductList

		GetIsVaildLangCodeMap(
			ctx context.Context,
		) map[int]string

		SetUpInMemory(
			ctx context.Context,
		)

		UpdateProduct(
			ctx context.Context,
			product domain.Product,
		)
	}

	// OrderUseCase ...
	OrderUseCase struct {
		ctx   context.Context
		order *domain.Order
	}

	// ToGateway ...
	ToGateway interface {
		GetAllProductList(
			ctx context.Context,
		) *domain.AllProductList

		UpdateProduct(
			ctx context.Context,
			product domain.Product,
		)
	}

	// ToPresenter ...
	ToPresenter interface {
		UpdateOrders(ctx context.Context)
		Shipment(ctx context.Context) error
	}
)

// NewUseCase ...
func NewUseCase(
	toDomain domain.ToDomain,
	toGateway ToGateway,
	toPresenter ToPresenter,
) ToUseCase {

	uscs := &useCase{
		ToDomain:    toDomain,
		ToGateway:   toGateway,
		ToPresenter: toPresenter,
	}

	return uscs
}
