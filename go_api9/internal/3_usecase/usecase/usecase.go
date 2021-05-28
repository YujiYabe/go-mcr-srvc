package usecase

import "context"

// Order ...
func (uc *UseCase) Order(ctx context.Context) error {
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

	// 返却
	err = uc.ToService.GetVegetables(ctx, requestVegetables)
	if err != nil {
		return err
	}

	return nil
}

// Dummy ...
func (uc *UseCase) Dummy(ctx context.Context) (string, error) {
	res, _ := uc.ToService.Dummy(ctx)

	return res, nil
}
