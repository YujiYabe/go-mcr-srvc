package monitor

import (
	"io"
	"sync"
	"text/template"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"

	"backend/internal/2_adapter/presenter"
	"backend/pkg"
)

var (
	myErr      *pkg.MyErr
	ordersChan = make(chan Orders)
)

func init() {
	myErr = pkg.NewMyErr("framework_driver", "monitor")
}

type (
	// Template ...
	Template struct {
		templates *template.Template
	}

	// Monitor ...
	Monitor struct {
		EchoEcho *echo.Echo
		Agents   map[string]*Agent
		Mutex    sync.RWMutex
	}

	// Orders ...
	Orders struct {
		Reserves  []string
		Assembles []string
		Completes []string
		Passes    []string
	}

	// Agent ...
	Agent struct {
		ID     string
		Socket *websocket.Conn
	}
)

var orders = &Orders{
	Reserves:  []string{},
	Assembles: []string{},
	Completes: []string{},
	Passes:    []string{},
}

// Render ...
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

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
