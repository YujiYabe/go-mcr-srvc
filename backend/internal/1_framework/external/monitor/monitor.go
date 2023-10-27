package monitor

import (
	"backend/internal/2_adapter/presenter"
)

// Monitor ...
type Monitor struct{}

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
