package register

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"

	"app/internal/2_adapter/controller"
	"app/internal/4_domain/domain"
)

var orderType domain.OrderType = "register"

const targetPath = "scripts/order/register"

type (
	// Register ...
	Register struct {
		Controller *controller.Controller
	}
)

// NewRegister ...
func NewRegister(ctrl *controller.Controller) *Register {
	rgstr := &Register{}
	rgstr.Controller = ctrl
	return rgstr
}

func (rgstr *Register) Start() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	rgstr.OrderAccept(targetPath)

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				switch {
				case event.Op&fsnotify.Create == fsnotify.Create:
					rgstr.OrderAccept(targetPath)

				case event.Op&fsnotify.Write == fsnotify.Write:
					rgstr.OrderAccept(targetPath)

				case event.Op&fsnotify.Remove == fsnotify.Remove:
				case event.Op&fsnotify.Rename == fsnotify.Rename:
				case event.Op&fsnotify.Chmod == fsnotify.Chmod:
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(targetPath)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func (rgstr *Register) OrderAccept(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fileName := file.Name()

		if strings.Count(fileName, ".") != 1 {
			continue
		}
		pos := strings.LastIndex(fileName, ".")
		if fileName[pos:] != ".json" {
			continue
		}

		path := filepath.Join(dir, file.Name())
		raw, err := ioutil.ReadFile(filepath.Clean(path))
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		order := &domain.Order{}
		err = json.Unmarshal(raw, order)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		ctx := context.Background()

		orderNumber, ctxValue := rgstr.Controller.Reserve(ctx, orderType)
		orderCtx := context.WithValue(ctx, orderNumber, ctxValue)
		go rgstr.Controller.Order(orderCtx, *order)

		newPath := strings.Replace(path, "json", orderNumber, 1)
		if err := os.Rename(path, newPath); err != nil {
			fmt.Println(err)
		}
	}

	return
}
