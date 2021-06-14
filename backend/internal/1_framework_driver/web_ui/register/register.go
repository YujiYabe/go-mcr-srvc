package register

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"

	"backend/internal/2_interface_adapter/controller"
	"backend/internal/4_enterprise_business_rule/entity"
	"backend/pkg"
)

var (
	orderType = "register"
	myErr     *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("infrastructure", "register")
}

type (
	// Register ...
	Register struct {
		Controller *controller.Controller
	}
)

// NewRegister ...
func NewRegister(ctrl *controller.Controller) *Register {
	rgstr := &Register{
		Controller: ctrl,
	}

	return rgstr
}

func (rgstr *Register) Start() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		myErr.Logging(err)
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
					rgstr.OrderAccept()

				case event.Op&fsnotify.Write == fsnotify.Write:
					rgstr.OrderAccept()

				case event.Op&fsnotify.Remove == fsnotify.Remove:
				case event.Op&fsnotify.Rename == fsnotify.Rename:
				case event.Op&fsnotify.Chmod == fsnotify.Chmod:
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				if err != nil {
					myErr.Logging(err)
				}
			}
		}
	}()

	err = watcher.Add(pkg.RegisterPath)
	if err != nil {
		myErr.Logging(err)
	}
	<-done
}

func (rgstr *Register) OrderAccept() {
	files, err := ioutil.ReadDir(pkg.RegisterPath)
	if err != nil {
		myErr.Logging(err)
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

		currentFilePath := filepath.Join(pkg.RegisterPath, currentFileName)

		raw, err := ioutil.ReadFile(filepath.Clean(currentFilePath))
		if err != nil {
			myErr.Logging(err)
			continue
		}

		product := &entity.Product{}
		err = json.Unmarshal(raw, product)
		if err != nil {
			myErr.Logging(err)
			continue
		}

		order := &entity.Order{
			Product: *product,
		}

		ctx := context.Background()

		rgstr.Controller.Reserve(ctx, order, orderType)
		newFileName := strings.Replace(currentFileName, "json", order.OrderInfo.OrderNumber, 1)
		newFilePath := filepath.Join(pkg.ReservedPath, newFileName)
		if err := os.Rename(currentFilePath, newFilePath); err != nil {
			myErr.Logging(err)
		}

		rgstr.Controller.Order(&ctx, order)
	}

	return
}
