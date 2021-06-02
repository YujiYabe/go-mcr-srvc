package usecase

import (
	"backend/internal/4_domain/domain"
	"context"
	"fmt"
	"sync"
)

// Order ...
func (uc *UseCase) Order(ctx context.Context, order *domain.Order) error {
	var err error
	fmt.Println("==============================")
	debugTarget := order.Product
	fmt.Printf("%#v\n", debugTarget)
	// fmt.Printf("%v\n", debugTarget)
	// fmt.Printf("%+v\n", debugTarget)
	// fmt.Printf("%T\n", debugTarget)
	fmt.Println("==============================")

	// オーダー解析
	assemble := uc.ToDomain.ParseOrder(ctx, order)

	// 材料取り出し
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		err = uc.ToService.GetVegetables(ctx, assemble.Vegetables)
	}()

	go func() {
		defer wg.Done()
		err = uc.ToService.GetPatties(ctx, assemble.Patties)
	}()

	go func() {
		defer wg.Done()
		err = uc.ToService.GetBans(ctx, assemble.Bans)
	}()

	go func() {
		defer wg.Done()
		err = uc.ToService.GetIngredients(ctx, assemble.Ingredients)
	}()

	wg.Wait()

	if err != nil {
		return err
	}

	// requestVegetables := map[string]int{"tomato": 1, "lettuce": 1}

	// err = uc.ToService.GetVegetables(ctx, requestVegetables)
	// if err != nil {
	// 	return err
	// }

	// // 調理
	// err = uc.ToService.(ctx, requestVegetables)
	// if err != nil {
	// 	return err
	// }

	// 出荷よー
	err = uc.ToService.Shipment(ctx, order)
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
