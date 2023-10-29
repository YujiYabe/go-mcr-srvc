package internal

import (
	"backend/internal/1_framework/db/sqlite"
	"backend/internal/1_framework/external/monitor"
	"backend/internal/1_framework/instacook"
	"backend/internal/2_adapter/controller"
)

type (
	app struct {
		instaCook *instacook.InstaCook
		monitor   *monitor.Monitor
	}
)

// NewApp ...
func NewApp() *app {
	isDemo := true
	// isDemo := false

	ctrl := controller.NewController(
		sqlite.NewToSQLite(),
		monitor.NewToMonitor(),
		isDemo,
	)

	a := &app{
		instaCook: instacook.NewInstaCook(ctrl),
		monitor:   monitor.NewMonitor(),
	}

	ctrl.Start()

	return a
}

// Start ...
func (receiver *app) Start() {
	receiver.instaCook.Start()
}
