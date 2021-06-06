package monitor

import (
	"backend/internal/2_adapter/service"
	"context"
	"io"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

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
	Assembles: []string{},
	Completes: []string{},
	Passes:    []string{},
}

var ordersChan = make(chan Orders)

// Render ...
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// NewMonitor ...
func NewMonitor() *Monitor {
	monitor := &Monitor{}
	monitor.EchoEcho = NewEcho()
	monitor.Agents = make(map[string]*Agent)

	return monitor
}

// NewToMonitor ...
func NewToMonitor() service.ToMonitor {
	s := new(Monitor)
	return s
}

// NewEcho ...
func NewEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	currentPath, _ := os.Getwd()
	WebPath := filepath.Join(currentPath, "web")
	IndexFilePath := filepath.Join(WebPath, "*.html")

	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob(IndexFilePath)),
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}__${status}__${method}__${uri}\n",
	}))

	e.Use(middleware.Recover())

	return e
}

// Start ...
func (monitor *Monitor) Start() {
	go monitor.Watching()
	go monitor.SendToAgents()

	currentPath, _ := os.Getwd()
	WebPath := filepath.Join(currentPath, "web")

	monitor.EchoEcho.Static("/web", WebPath)

	monitor.EchoEcho.GET("/", monitor.Index)
	monitor.EchoEcho.GET("/ws", monitor.WebSocket)
	monitor.EchoEcho.Logger.Fatal(monitor.EchoEcho.Start(":4567"))
}

// UpdateOrders ...
func (monitor *Monitor) UpdateOrders(ctx context.Context, orderNumber string, phase string) error {

	switch phase {
	case "assemble":
		orders.Assembles = append(orders.Assembles, orderNumber)
	case "complete":
		orders.Assembles = remove(orders.Assembles, orderNumber)
		orders.Completes = append(orders.Completes, orderNumber)
	case "pass":
		orders.Completes = remove(orders.Completes, orderNumber)
		orders.Passes = append(orders.Passes, orderNumber)
	}

	ordersChan <- *orders

	return nil
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
