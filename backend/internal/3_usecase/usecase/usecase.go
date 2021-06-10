package usecase

import (
	"context"
	"fmt"
	"sync"
	"time"

	"backend/internal/4_domain/domain"
	"backend/pkg"
)

var (
	myErr        *pkg.MyErr
	orderUsecase = make(chan OrderUsecase)
)

func init() {
	myErr = pkg.NewMyErr("usecase", "usecase")
}

// Start ...
func (uc *UseCase) Start() {
	go uc.bulkOrder()
}

// Reserve ...
func (uc *UseCase) Reserve(ctx context.Context, orderinfo *domain.OrderInfo) {
	uc.ToService.UpdateOrders(ctx, orderinfo.OrderNumber, "reserve")

	return
}

// Order ...
func (uc *UseCase) Order(ctx *context.Context, order *domain.Order) error {
	ou := &OrderUsecase{
		ctx:   ctx,
		order: order,
	}

	orderUsecase <- *ou
	fmt.Println(len(orderUsecase))
	fmt.Println(cap(orderUsecase))
	return nil
}

func (uc *UseCase) bulkOrder() {
	var err error

	for {
		ou := <-orderUsecase

		// go func() {
		ctxWithTimeout, _ := context.WithTimeout(*ou.ctx, time.Minute*10)

		uc.ToService.UpdateOrders(ctxWithTimeout, ou.order.OrderInfo.OrderNumber, "assemble")

		// オーダー解析
		assemble := uc.ToDomain.ParseOrder(ctxWithTimeout, ou.order)

		// 材料取り出し
		err = uc.getFoodstuff(ctxWithTimeout, assemble)
		if err != nil {
			myErr.Logging(err)
		}

		// 調理
		err = uc.cookFoodstuff(ctxWithTimeout, ou.order, assemble)
		if err != nil {
			myErr.Logging(err)
		}

		// 出荷よー
		err = uc.ToService.Shipment(ctxWithTimeout, ou.order)
		if err != nil {
			myErr.Logging(err)
		}

		uc.ToService.UpdateOrders(ctxWithTimeout, ou.order.OrderInfo.OrderNumber, "complete")
		// }()
	}

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
