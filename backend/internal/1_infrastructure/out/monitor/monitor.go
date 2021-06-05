package monitor

import (
	"backend/internal/2_adapter/service"
	"backend/internal/4_domain/domain"
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
		// Member     []string
		// Controller *controller.Controller
		Orders Orders
	}

	// Orders ...
	Orders struct {
		Assembled []string
		Completed []string
		Passed    []string
	}

	// Agent ...
	Agent struct {
		ID     string
		Socket *websocket.Conn
	}
)

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

	currentPath, _ := os.Getwd()
	WebPath := filepath.Join(currentPath, "web")

	monitor.EchoEcho.Static("/web", WebPath)

	monitor.EchoEcho.GET("/", monitor.Index)
	monitor.EchoEcho.GET("/ws", monitor.WebSocket)
	monitor.EchoEcho.Logger.Fatal(monitor.EchoEcho.Start(":4567"))
}

// UpdateOrders ...
func (monitor *Monitor) UpdateOrders(context.Context, *domain.Order) error {

	return nil
}
