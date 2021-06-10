package usecase

import (
	"context"
	"sync"

	"backend/internal/4_domain/domain"
	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("usecase", "usecase")
}

// Reserve ...
func (uc *UseCase) Reserve(ctx context.Context, orderinfo *domain.OrderInfo) error {
	// var err error

	uc.ToService.UpdateOrders(ctx, orderinfo.OrderNumber, "reserve")

	return nil
}

// Order ...
func (uc *UseCase) Order(ctx context.Context, order *domain.Order) error {
	var err error

	uc.ToService.UpdateOrders(ctx, order.OrderInfo.OrderNumber, "assemble")

	// オーダー解析
	assemble := uc.ToDomain.ParseOrder(ctx, order)

	// 材料取り出し
	err = uc.getFoodstuff(ctx, assemble)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	// 調理
	err = uc.cookFoodstuff(ctx, order, assemble)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	// 出荷よー
	err = uc.ToService.Shipment(ctx, order)
	if err != nil {
		myErr.Logging(err)
		return err
	}
	uc.ToService.UpdateOrders(ctx, order.OrderInfo.OrderNumber, "complete")

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
		myErr.Logging(err)
		return err
	}

	return nil
}

func (uc *UseCase) cookFoodstuff(ctx context.Context, order *domain.Order, assemble *domain.Assemble) error {
	if len(order.Product.Hamburgers) > 0 {
		err := uc.ToDomain.CookHamburgers(ctx, order.Product.Hamburgers)
		if err != nil {
			myErr.Logging(err)
			return err
		}
	}

	return nil
}
