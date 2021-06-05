package monitor

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

const targetPath = "yummy"

func (monitor *Monitor) Watching() {
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
				case event.Op&fsnotify.Write == fsnotify.Write:
				case event.Op&fsnotify.Remove == fsnotify.Remove:
					monitor.Passed(targetPath)

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

func (monitor *Monitor) Passed(currentDir string) {

	return
}
