package monitor

import (
	"net/http"
	"os"
	"text/template"
)

type Monitor struct{}
type (
	// ServerInfo ...
	ServerInfo struct {
		Address string
		Port    string
	}
)

// NewMonitor ...
func NewMonitor() *Monitor {
	mntr := &Monitor{}
	return mntr
}

// Start ...
func (mb *Monitor) Start() {
	dir, _ := os.Getwd()
	http.Handle("/web/css/", http.StripPrefix("/web/css/", http.FileServer(http.Dir(dir+"/web/css/"))))
	http.Handle("/web/js/", http.StripPrefix("/web/js/", http.FileServer(http.Dir(dir+"/web/js/"))))
	http.Handle("/web/vue/", http.StripPrefix("/web/vue/", http.FileServer(http.Dir(dir+"/web/vue/"))))

	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":4567", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/index.html") // ParseFilesを使う
	if err != nil {
		panic(err)
	}

	serverInfo := &ServerInfo{
		Port: "4567",
	}

	err = tmpl.Execute(w, serverInfo)
	if err != nil {
		panic(err)
	}
}
