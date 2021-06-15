package usecase

import (
	"context"
	"sync"
	"time"

	"github.com/anikhasibul/queue"

	"backend/internal/4_enterprise_business_rule/entity"
	"backend/pkg"
)

var (
	myErr        *pkg.MyErr
	orderUsecase = make(chan OrderUsecase)
)

func init() {
	myErr = pkg.NewMyErr("application_business_rule", "usecase")
}

// Start ...
func (uc *useCase) Start() {
	go uc.bulkOrder()
}

// Reserve ...
func (uc *useCase) Reserve(ctx context.Context, orderinfo *entity.OrderInfo) {
	uc.ToPresenter.UpdateOrders(ctx, orderinfo.OrderNumber, "reserve")
}

// Order ...
func (uc *useCase) Order(ctx *context.Context, order *entity.Order) error {
	ou := &OrderUsecase{
		ctx:   ctx,
		order: order,
	}

	orderUsecase <- *ou

	return nil
}

func (uc *useCase) getFoodstuff(ctx context.Context, assemble *entity.Assemble) error {
	var err error
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = uc.ToGateway.GetVegetables(ctx, assemble.Vegetables)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = uc.ToGateway.GetPatties(ctx, assemble.Patties)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = uc.ToGateway.GetBans(ctx, assemble.Bans)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = uc.ToGateway.GetIngredients(ctx, assemble.Ingredients)
	}()

	wg.Wait()
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}

func (uc *useCase) cookFoodstuff(ctx context.Context, order *entity.Order, assemble *entity.Assemble) error {
	if len(order.Product.Hamburgers) > 0 {
		err := uc.ToEntity.CookHamburgers(ctx, order.Product.Hamburgers)
		if err != nil {
			myErr.Logging(err)
			return err
		}
	}

	return nil
}

func (uc *useCase) bulkOrder() {
	var err error
	q := queue.New(pkg.AssembleNumber)
	defer q.Close()

	for {
		ou := <-orderUsecase
		q.Add()
		go func() {
			defer q.Done()
			ctxWithTimeout, _ := context.WithTimeout(*ou.ctx, time.Minute*10)

			uc.ToPresenter.UpdateOrders(ctxWithTimeout, ou.order.OrderInfo.OrderNumber, "assemble")

			// オーダー解析
			assemble := uc.ToEntity.ParseOrder(ctxWithTimeout, ou.order)

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
			err = uc.ToPresenter.Shipment(ctxWithTimeout, ou.order)
			if err != nil {
				myErr.Logging(err)
			}

			uc.ToPresenter.UpdateOrders(ctxWithTimeout, ou.order.OrderInfo.OrderNumber, "complete")
		}()
	}

}
