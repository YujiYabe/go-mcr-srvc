package usecase

import (
	"context"
	"sync"

	domain "backend/internal/4_domain"

	"github.com/anikhasibul/queue"

	"backend/pkg"
)

var (
	myErr        *pkg.MyErr
	orderUseCase = make(chan OrderUseCase)
)

func init() {
	myErr = pkg.NewMyErr("application_business_rule", "usecase")
}

// Start ...
func (uscs *useCase) Start() {
	go uscs.bulkOrder()
}

// Reserve ...
func (uscs *useCase) Reserve(ctx context.Context, orderInfo *domain.OrderInfo) {
	uscs.ToPresenter.UpdateOrders(ctx, orderInfo.OrderNumber, "reserve") // オーダー情報更新
}

// Order ...
func (uscs *useCase) Order(ctx *context.Context, order *domain.Order) error {
	ou := &OrderUseCase{
		ctx:   ctx,
		order: order,
	}

	orderUseCase <- *ou

	return nil
}

func (uscs *useCase) bulkOrder() {
	var err error

	q := queue.New(pkg.AssembleNumber) // 擬似的に同時進行できるキャパシティを設定
	defer q.Close()

	for {
		ou := <-orderUseCase
		q.Add()
		go func() {
			defer q.Done()

			// オーダー情報更新
			uscs.ToPresenter.UpdateOrders(*ou.ctx, ou.order.OrderInfo.OrderNumber, "assemble")

			// オーダー解析
			assemble := uscs.ToDomain.ParseOrder(*ou.ctx, ou.order)

			// 材料取り出し
			err = uscs.getFoodstuff(*ou.ctx, assemble)
			if err != nil {
				myErr.Logging(err)
			}

			// 調理
			err = uscs.cookFoodstuff(*ou.ctx, ou.order, assemble)
			if err != nil {
				myErr.Logging(err)
			}

			// 出荷よー
			err = uscs.ToPresenter.Shipment(*ou.ctx, ou.order)
			if err != nil {
				myErr.Logging(err)
			}

			uscs.ToPresenter.UpdateOrders(*ou.ctx, ou.order.OrderInfo.OrderNumber, "complete")
		}()
	}

}

func (uscs *useCase) getFoodstuff(ctx context.Context, assemble *domain.Assemble) error {
	var err error
	var wg sync.WaitGroup

	// 材料取り出し
	// 同時進行処理
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = uscs.ToGateway.GetVegetables(ctx, assemble.Vegetables)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = uscs.ToGateway.GetIngredients(ctx, assemble.Ingredients)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = uscs.ToGateway.GetPatties(ctx, assemble.Patties)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err = uscs.ToGateway.GetBans(ctx, assemble.Bans)
	}()

	wg.Wait()
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}

func (uscs *useCase) cookFoodstuff(ctx context.Context, order *domain.Order, assemble *domain.Assemble) error {
	// オーダーにハンバーガーが含まれていれば調理
	if len(order.Product.Hamburgers) > 0 {
		err := uscs.ToDomain.CookHamburgers(ctx, order.Product.Hamburgers)
		if err != nil {
			myErr.Logging(err)
			return err
		}
	}

	return nil
}
