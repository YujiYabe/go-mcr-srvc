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
	instaCook *web.InstaCook
	monitor   *monitor.Monitor
}

// NewApp ...
func NewApp() *app {

	ctrl := controller.NewController(
		sqlite.NewToSQLite(),
		monitor.NewToMonitor(),
		isDemo,
	)

	a := &app{
		instaCook: web.NewInstaCook(ctrl),
		monitor:   monitor.NewMonitor(),
	}

	ctrl.Start()

	return a
}

// Start ...
func (receiver *app) Start() {
	receiver.instaCook.Start(isShowRoute)
}
