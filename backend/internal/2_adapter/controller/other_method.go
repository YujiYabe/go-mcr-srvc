package controller

import (
	"context"

	domain "backend/internal/4_domain"
)

// Start ...
func (receiver *controller) Start() {
	ctx := context.Background()

	receiver.ToUseCase.Start(ctx) // 初期処理 DBをインメモリに保存
}

// allergy -----------------------
// GetAllergyDefault ...
func (receiver *controller) GetAllergyDefault(
	ctx context.Context,
) domain.Allergy {
	return receiver.ToUseCase.GetAllergyDefault(
		ctx,
	)
}

// GetAllergyList ...
func (receiver *controller) GetAllergyList(
	ctx context.Context,
) domain.AllergyList {
	return receiver.ToUseCase.GetAllergyList(
		ctx,
	)
}

// language -----------------------
// GetIsVaildLangCodeMap ...
func (receiver *controller) GetIsVaildLangCodeMap(
	ctx context.Context,
) map[int]string {
	return receiver.ToUseCase.GetIsVaildLangCodeMap(
		ctx,
	)
}

// websocket -----------------------
// DistributeOrder ...
func (receiver *controller) DistributeOrder(
	ctx context.Context,
) {
	receiver.ToUseCase.DistributeOrder(
		ctx,
	)
}
