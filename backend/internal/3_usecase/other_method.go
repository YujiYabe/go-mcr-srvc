package usecase

import (
	"context"

	domain "backend/internal/4_domain"
)

// Start ...
func (receiver *useCase) Start(ctx context.Context) {
	receiver.SetUpInMemory(ctx)
}

// allergy -----------------------
// GetAllergyDefault ...
func (receiver *useCase) GetAllergyDefault(
	ctx context.Context,
) domain.Allergy {
	return receiver.ToDomain.GetAllergyDefault(
		ctx,
	)
}

// GetAllergyList ...
func (receiver *useCase) GetAllergyList(
	ctx context.Context,
) domain.AllergyList {
	return receiver.ToDomain.GetAllergyList(
		ctx,
	)
}

// language -----------------------
// GetIsVaildLangCodeMap ...
func (receiver *useCase) GetIsVaildLangCodeMap(
	ctx context.Context,
) map[int]string {
	return receiver.ToDomain.GetIsVaildLangCodeMap()
}

// websocket -----------------------
// DistributeOrder ...
func (receiver *useCase) DistributeOrder(
	ctx context.Context,
) {
	orderList := receiver.ToDomain.GetOrderList(
		ctx,
	)

	receiver.ToPresenter.DistributeOrder(
		ctx,
		&orderList,
	)

}
