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
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// NewMonitor ...
func NewMonitor() *Monitor {
	mntr := &Monitor{}
	mntr.EchoEcho = NewEcho()
	mntr.Agents = make(map[string]*Agent)

	return mntr
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
func (mntr *Monitor) Start() {
	mntr.RemoveYummy()

	go mntr.Watching()
	go mntr.SendToAgents()

	mntr.EchoEcho.Renderer = &Template{
		templates: template.Must(template.ParseGlob(pkg.IndexPath)),
	}

	// mntr.EchoEcho.Static("/web", pkg.WebPath)

	mntr.EchoEcho.GET("/", mntr.Index)
	mntr.EchoEcho.GET("/ws", mntr.WebSocket)
	mntr.EchoEcho.Logger.Fatal(mntr.EchoEcho.Start(":" + pkg.MonitorPort))
}

// UpdateOrders ...
func (mntr *Monitor) UpdateOrders(ctx context.Context) {
	return
}

func remove() error {
	return nil
}
