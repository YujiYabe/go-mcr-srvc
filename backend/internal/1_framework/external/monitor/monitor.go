package monitor

import (
	"backend/internal/2_adapter/presenter"
	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("framework_driver", "monitor")
}

type (

	// Monitor ...
	Monitor struct {
	}
)

// NewMonitor ...
func NewMonitor() *Monitor {
	mntr := &Monitor{}

	return mntr
}

// NewToMonitor ...
func NewToMonitor() presenter.ToMonitor {
	s := new(Monitor)
	return s
}
