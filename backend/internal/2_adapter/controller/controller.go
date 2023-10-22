package controller

import (
	"context"

	"backend/internal/2_adapter/gateway"
	"backend/internal/2_adapter/presenter"
	usecase "backend/internal/3_usecase"
	domain "backend/internal/4_domain"

	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("interface_adapter", "controller")
}

type (
	// controller ...
	controller struct {
		UseCase     usecase.ToUseCase
		OrderNumber int
	}

	// orderChannel ...
	orderChannel struct {
		order *domain.Order
	}

	// ToController ...
	ToController interface {
		Start()
		Reserve(ctx context.Context, order *domain.Order, orderType string)
		Order(ctx *context.Context, order *domain.Order)
	}
)

// orderChannel ...
var odrChnnl = make(chan orderChannel)

// NewController ...
func NewController(
	toSqlite gateway.ToSqlite,
	// toPostgres gateway.ToPostgres,
	// toMysql gateway.ToMysql,
	// toMongo gateway.ToMongo,
	toShipment presenter.ToShipment,
	toMonitor presenter.ToMonitor,
) ToController {
	toDomain := domain.NewDomain()
	toGateway := gateway.NewGateway(
		toSqlite,
		// toPostgres,
		// toMysql,
		// toMongo,
	)
	toPresenter := presenter.NewPresenter(
		toShipment,
		toMonitor,
	)
	uscs := usecase.NewUseCase(
		toDomain,
		toGateway,
		toPresenter,
	)
	ct := &controller{
		UseCase: uscs,
	}

	return ct
}

func (ctrl *controller) Start() {
	go ctrl.bulkOrder()
	go ctrl.UseCase.Start()
}

// Reserve ...
func (ctrl *controller) Reserve(
	ctx context.Context,
	order *domain.Order,
	orderType string,
) {
	ctrl.UseCase.Reserve(ctx) // オーダー情報更新
}

// Order ...
func (ctrl *controller) Order(
	ctx *context.Context,
	order *domain.Order,
) {
	oc := &orderChannel{
		order: order,
	}

	// オーダー番号をweb_uiに即時返却する必要があるため、
	// 後続処理をチャネル経由で処理
	odrChnnl <- *oc
}

func (ctrl *controller) bulkOrder() {
}
