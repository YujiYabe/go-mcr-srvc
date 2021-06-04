package monitor

import (
	"net/http"
	"text/template"
)

type Monitor struct{}
type Page struct {
	Title string
	Count int
}

// NewMonitor ...
func NewMonitor() *Monitor {
	mntr := &Monitor{}
	return mntr
}

// Start ...
func (mb *Monitor) Start() {
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":4567", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	page := Page{"Hello World.", 1}
	tmpl, err := template.ParseFiles("web/layout.html") // ParseFilesを使う
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}
