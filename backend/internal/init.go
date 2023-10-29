package internal

import (
	"backend/internal/1_framework/db/sqlite"
	"backend/internal/1_framework/external/monitor"
	"backend/internal/1_framework/web"
	"backend/internal/2_adapter/controller"
)

var (
	isDemo = true
	// isDemo = false
	// isShowRoute = true
	isShowRoute = false
)

type app struct {
	web     *web.Web
	monitor *monitor.Monitor
}

// NewApp ...
func NewApp() *app {

	ctrl := controller.NewController(
		sqlite.NewToSQLite(),
		monitor.NewToMonitor(),
		isDemo,
	)

	a := &app{
		web:     web.NewWeb(ctrl),
		monitor: monitor.NewMonitor(),
	}

	ctrl.Start() // 初期処理 DBをインメモリに保存

	return a
}

// Start ...
func (receiver *app) Start() {
	receiver.web.Start(isShowRoute)
}
