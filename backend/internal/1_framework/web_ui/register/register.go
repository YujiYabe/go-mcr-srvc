package register

import (
	"log"

	"github.com/fsnotify/fsnotify"

	"backend/internal/2_adapter/controller"
	"backend/pkg"
)

var (
	orderType = "register"
	myErr     *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("framework_driver", "register")
}

type (
	// Register ...
	Register struct {
		Controller controller.ToController
	}
)

// NewRegister ...
func NewRegister(ctrl controller.ToController) *Register {
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

}
