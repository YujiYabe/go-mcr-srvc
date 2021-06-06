package usecase

import (
	"backend/internal/4_domain/domain"
	"context"
	"sync"
)

// Order ...
func (uc *UseCase) Order(ctx context.Context, order *domain.Order) error {
	var err error

	err = uc.ToService.UpdateOrders(ctx, order.OrderInfo.OrderNumber, "assemble")
	if err != nil {
		return err
	}

	// オーダー解析
	assemble := uc.ToDomain.ParseOrder(ctx, order)

	// 材料取り出し
	err = uc.getFoodstuff(ctx, assemble)
	if err != nil {
		return err
	}

	// 調理
	err = uc.cookFoodstuff(ctx, order, assemble)
	if err != nil {
		return err
	}
	// limit := 3
	// slots := make(chan struct{}, limit)
	// var wg sync.WaitGroup
	// for {
	// 	slots <- struct{}{}
	// 	wg.Add(1)
	// 	go func() {
	// 		err := uc.cookFoodstuff(ctx, order, assemble)
	// 		if err != nil {
	// 			fmt.Printf("%+v\n\n", err)
	// 		}
	// 		<-slots
	// 		wg.Done()
	// 	}()
	// }
	// wg.Wait()

	// 出荷よー
	err = uc.ToService.Shipment(ctx, order)
	if err != nil {
		return err
	}

	err = uc.ToService.UpdateOrders(ctx, order.OrderInfo.OrderNumber, "complete")

	return nil
}

func (uc *UseCase) getFoodstuff(ctx context.Context, assemble *domain.Assemble) error {
	var err error
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = uc.ToService.GetVegetables(ctx, assemble.Vegetables)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = uc.ToService.GetPatties(ctx, assemble.Patties)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = uc.ToService.GetBans(ctx, assemble.Bans)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = uc.ToService.GetIngredients(ctx, assemble.Ingredients)
	}()

	wg.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) cookFoodstuff(ctx context.Context, order *domain.Order, assemble *domain.Assemble) error {
	var err error
	if len(order.Product.Hamburgers) > 0 {
		err = uc.ToDomain.CookHamburgers(ctx, order.Product.Hamburgers)
	}

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
