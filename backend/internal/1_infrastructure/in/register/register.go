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

	"backend/internal/2_adapter/controller"
	"backend/internal/4_domain/domain"

	"github.com/fsnotify/fsnotify"
)

var orderType = "register"

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

func (rgstr *Register) OrderAccept(currentDir string) {
	files, err := ioutil.ReadDir(currentDir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		currentFileName := file.Name()

		if strings.Count(currentFileName, ".") != 1 {
			continue
		}
		pos := strings.LastIndex(currentFileName, ".")
		if currentFileName[pos:] != ".json" {
			continue
		}

		currentFilePath := filepath.Join(currentDir, file.Name())
		newDir := filepath.Join(currentDir, "reserved")

		raw, err := ioutil.ReadFile(filepath.Clean(currentFilePath))
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		order := &domain.Order{}
		product := &domain.Product{}
		err = json.Unmarshal(raw, product)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		order.Product = *product

		ctx := context.Background()

		rgstr.Controller.Reserve(ctx, order, orderType)

		go rgstr.Controller.Order(ctx, order)

		newFileName := strings.Replace(currentFileName, "json", order.OrderInfo.OrderNumber, 1)
		newFilePath := filepath.Join(newDir, newFileName)
		if err := os.Rename(currentFilePath, newFilePath); err != nil {
			fmt.Println(err)
		}
	}

	return
}
