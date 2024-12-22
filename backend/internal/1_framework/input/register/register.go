package register

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"

	"backend/internal/2_adapter/controller"
	domain "backend/internal/4_domain"
	"backend/pkg"
)

type (
	// Register ...
	Register struct {
		Controller controller.ToController
	}
)

// NewRegister ...
func NewRegister(receiver controller.ToController) *Register {
	register := &Register{
		Controller: receiver,
	}

	return register
}

func (receiver *Register) Start() {
	ctx := context.Background()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		pkg.Logging(ctx, err)
		log.Fatal(err)
	}
	defer watcher.Close()
	done := make(chan bool)

	// go func() {
	// 	for {
	// 		select {
	// 		case event, ok := <-watcher.Events:
	// 			if !ok {
	// 				return
	// 			}
	// 			switch {
	// 			case event.Op&fsnotify.Create == fsnotify.Create:
	// 				receiver.OrderAccept()

	// 			case event.Op&fsnotify.Write == fsnotify.Write:
	// 				receiver.OrderAccept()

	// 			case event.Op&fsnotify.Remove == fsnotify.Remove:
	// 			case event.Op&fsnotify.Rename == fsnotify.Rename:
	// 			case event.Op&fsnotify.Chmod == fsnotify.Chmod:
	// 			}

	// 		case err, ok := <-watcher.Errors:
	// 			if !ok {
	// 				return
	// 			}
	// 			if err != nil {
	// 				pkg.Logging(ctx, err)
	// 			}
	// 		}
	// 	}
	// }()

	err = watcher.Add(pkg.RegisterPath)
	if err != nil {
		pkg.Logging(ctx, err)
	}
	<-done
}

func (receiver *Register) OrderAccept() {
	ctx := context.Background()
	files, err := os.ReadDir(pkg.RegisterPath)
	if err != nil {
		pkg.Logging(ctx, err)
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

		raw, err := os.ReadFile(filepath.Clean(currentFilePath))
		if err != nil {
			pkg.Logging(ctx, err)
			continue
		}

		product := &domain.Product{}
		err = json.Unmarshal(raw, product)
		if err != nil {
			pkg.Logging(ctx, err)
			continue
		}

		order := &domain.Order{
			Product: *product,
		}

		// 標準コンテキストを取得
		ctx := context.Background()

		// receiver.Controller.Reserve(ctx, order, orderType) // オーダー番号発行
		// receiver.Controller.Order(&ctx, order)             // オーダー

		newExtension := order.OrderInfo.OrderNumber + ".json"
		newFileName := strings.Replace(currentFileName, "json", newExtension, 1)

		newFilePath := filepath.Join(pkg.ReservedPath, newFileName)
		err = os.Rename(currentFilePath, newFilePath) // オーダー番号返却

		if err != nil {
			pkg.Logging(ctx, err)
		}
	}

}
