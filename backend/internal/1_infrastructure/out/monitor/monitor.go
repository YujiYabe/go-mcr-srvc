package monitor

import (
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
	mntr := &Monitor{}
	mntr.EchoEcho = NewEcho()
	mntr.Agents = make(map[string]*Agent)

	return mntr
}

// NewEcho ...
func NewEcho() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	// 	http.Handle("/web/css/", http.StripPrefix("/web/css/", http.FileServer(http.Dir(dir+"/web/css/"))))
	// http.Handle("/web/js/", http.StripPrefix("/web/js/", http.FileServer(http.Dir(dir+"/web/js/"))))
	// http.Handle("/web/vue/", http.StripPrefix("/web/vue/", http.FileServer(http.Dir(dir+"/web/vue/"))))

	// http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	// tmpl, err := template.ParseFiles("web/index.html") // ParseFilesを使う
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
func (mntr *Monitor) Start() {
	currentPath, _ := os.Getwd()
	WebPath := filepath.Join(currentPath, "web")
	mntr.EchoEcho.Static("/web", WebPath)

	mntr.EchoEcho.GET("/", mntr.Index)
	mntr.EchoEcho.GET("/ws", mntr.WebSocket)
	mntr.EchoEcho.Logger.Fatal(mntr.EchoEcho.Start(":4567"))
}
