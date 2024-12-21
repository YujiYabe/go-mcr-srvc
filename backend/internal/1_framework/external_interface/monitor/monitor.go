package monitor

import (
	"context"
	"io"
	"sync"
	"text/template"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

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
func (receiver *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return receiver.templates.ExecuteTemplate(w, name, data)
}

// NewMonitor ...
func NewMonitor() *Monitor {
	monitor := &Monitor{}
	monitor.EchoEcho = NewEcho()
	monitor.Agents = make(map[string]*Agent)

	return monitor
}

// NewToMonitor ...
func NewToMonitor() presenter.ToMonitor {
	s := new(Monitor)
	return s
}

// NewEcho ...
func NewEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}__${status}__${method}__${uri}\n",
	}))

	e.Use(middleware.Recover())

	return e
}

// Start ...
func (receiver *Monitor) Start() {
	receiver.RemoveYummy()

	go receiver.Watching()
	go receiver.SendToAgents()

	receiver.EchoEcho.Renderer = &Template{
		templates: template.Must(template.ParseGlob(pkg.IndexPath)),
	}

	receiver.EchoEcho.Static("/web", pkg.WebPath)

	receiver.EchoEcho.GET("/", receiver.Index)
	receiver.EchoEcho.GET("/ws", receiver.WebSocket)
	receiver.EchoEcho.Logger.Fatal(receiver.EchoEcho.Start(":" + pkg.MonitorPort))
}

// UpdateOrders ...
func (receiver *Monitor) UpdateOrders(ctx context.Context, orderNumber string, phase string) {

	switch phase {
	case "reserve":
		orders.Reserves = append(orders.Reserves, orderNumber)
	case "assemble":
		orders.Reserves = remove(orders.Reserves, orderNumber)
		orders.Assembles = append(orders.Assembles, orderNumber)
	case "complete":
		orders.Assembles = remove(orders.Assembles, orderNumber)
		orders.Completes = append(orders.Completes, orderNumber)
	case "pass":
		orders.Completes = remove(orders.Completes, orderNumber)
		orders.Passes = append(orders.Passes, orderNumber)
	}

	ordersChan <- *orders
}

func remove(strings []string, search string) []string {
	result := []string{}
	for _, v := range strings {
		if v != search {
			result = append(result, v)
		}
	}
	return result
}
