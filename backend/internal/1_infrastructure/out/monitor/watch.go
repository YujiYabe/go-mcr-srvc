package monitor

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const targetPath = "yummy"

func (monitor *Monitor) Watching() {
	currentPath, _ := os.Getwd()
	yummyPath := filepath.Join(currentPath, targetPath)
	// var err error
	var currentfiles []string

	for {
		files, err := ioutil.ReadDir(yummyPath)
		if err != nil {
			panic(err)
		}

		var newFiles []string
		for _, file := range files {
			newFiles = append(newFiles, file.Name())
		}

		monitor.passedCheck(currentfiles, newFiles)

		currentfiles = newFiles

		time.Sleep(1 * time.Second)
	}
}

func (monitor *Monitor) passedCheck(currentfiles, newFiles []string) []string {
	//最新のリストからファイルが削除されていれば渡しずみ判断
	var passedFiles []string
	for _, currentfile := range currentfiles {
		isExist := false
		for _, newFile := range newFiles {
			if currentfile == newFile {
				isExist = true
				continue
			}
		}
		if !isExist {
			monitor.Orders.Completes = remove(monitor.Orders.Completes, currentfile)
			monitor.Orders.Passes = append(monitor.Orders.Passes, currentfile)
		}
	}
	return passedFiles
}
