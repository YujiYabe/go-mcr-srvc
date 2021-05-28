package usecase

import (
	"app/internal/4_domain/domain"
	"context"
	"fmt"
)

// Order ...
func (uc *UseCase) Order(ctx context.Context, order domain.Order) error {

	fmt.Println("==============================")
	debugTarget := order
	fmt.Printf("%#v\n", debugTarget)
	fmt.Println("==============================")

	requestVegetables := map[string]int{"tomato": 1, "lettuce": 1}

	// 材料取り出し
	err := uc.ToService.GetVegetables(ctx, requestVegetables)
	if err != nil {
		return err
	}

	// // 調理
	// err = uc.ToService.(ctx, requestVegetables)
	// if err != nil {
	// 	return err
	// }

	// // 返却
	// err = uc.ToService.GetVegetables(ctx, requestVegetables)
	// if err != nil {
	// 	return err
	// }

	return nil
}

// Dummy ...
func (uc *UseCase) Dummy(ctx context.Context) (string, error) {
	res, _ := uc.ToService.Dummy(ctx)

	return res, nil
}
