package monitor

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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
			ctx := context.Background()
			monitor.UpdateOrders(ctx, strings.TrimRight(currentfile, ".json"), "pass")
		}
	}
	return passedFiles
}
