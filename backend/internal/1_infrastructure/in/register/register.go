package register

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"

	"app/internal/2_adapter/controller"
	"app/internal/4_domain/domain"
)

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

func (rgstr *Register) OrderAccept(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	for _, path := range paths {
		raw, err := ioutil.ReadFile(path)
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
		rgstr.Controller.Order(ctx, *order)

	}

	return paths
}
