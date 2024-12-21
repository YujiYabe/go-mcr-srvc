package usecase

import (
	"context"
	"sync"

	"github.com/anikhasibul/queue"

	domain "backend/internal/4_domain"
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
func (receiver *useCase) Start() {
	go receiver.bulkOrder()
}

// Reserve ...
func (receiver *useCase) Reserve(ctx context.Context, orderInfo *domain.OrderInfo) {
	receiver.ToPresenter.UpdateOrders(ctx, orderInfo.OrderNumber, "reserve") // オーダー情報更新
}

// Order ...
func (receiver *useCase) Order(ctx *context.Context, order *domain.Order) error {
	ou := &OrderUsecase{
		ctx:   ctx,
		order: order,
	}

	orderUsecase <- *ou

	return nil
}

func (receiver *useCase) bulkOrder() {
	var err error

	q := queue.New(pkg.AssembleNumber) // 擬似的に同時進行できるキャパシティを設定
	defer q.Close()

	for {
		ou := <-orderUsecase
		q.Add()
		go func() {
			defer q.Done()

			// オーダー情報更新
			receiver.ToPresenter.UpdateOrders(*ou.ctx, ou.order.OrderInfo.OrderNumber, "assemble")

			// オーダー解析
			assemble := receiver.ToDomain.ParseOrder(*ou.ctx, ou.order)

			// 材料取り出し
			err = receiver.getFoodstuff(*ou.ctx, assemble)
			if err != nil {
				myErr.Logging(err)
			}

			// 調理
			err = receiver.cookFoodstuff(*ou.ctx, ou.order, assemble)
			if err != nil {
				myErr.Logging(err)
			}

			// 出荷よー
			err = receiver.ToPresenter.Shipment(*ou.ctx, ou.order)
			if err != nil {
				myErr.Logging(err)
			}

			receiver.ToPresenter.UpdateOrders(*ou.ctx, ou.order.OrderInfo.OrderNumber, "complete")
		}()
	}

}

func (receiver *useCase) getFoodstuff(ctx context.Context, assemble *domain.Assemble) error {
	var err error
	var wg sync.WaitGroup

	// 材料取り出し
	// 同時進行処理
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = receiver.ToGateway.GetVegetables(ctx, assemble.Vegetables)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = receiver.ToGateway.GetIngredients(ctx, assemble.Ingredients)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = receiver.ToGateway.GetPatties(ctx, assemble.Patties)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = receiver.ToGateway.GetBans(ctx, assemble.Bans)
	}()

	wg.Wait()
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}

func (receiver *useCase) cookFoodstuff(ctx context.Context, order *domain.Order, _ *domain.Assemble) error {
	// オーダーにハンバーガーが含まれていれば調理
	if len(order.Product.Hamburgers) > 0 {
		err := receiver.ToDomain.CookHamburgers(ctx, order.Product.Hamburgers)
		if err != nil {
			myErr.Logging(err)
			return err
		}
	}

	return nil
}
