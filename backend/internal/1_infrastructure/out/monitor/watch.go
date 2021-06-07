package monitor

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"backend/pkg"
)

func (monitor *Monitor) RemoveYummy() {
	yummyFiles, err := ioutil.ReadDir(pkg.YummyPath)
	if err != nil {
		myErr.Logging(err)
	}

	for _, yummyFile := range yummyFiles {
		err := os.Remove(filepath.Join(pkg.YummyPath, yummyFile.Name()))
		if err != nil {
			myErr.Logging(err)
		}
	}

}

func (monitor *Monitor) Watching() {
	var currentfiles []string

	for {
		files, err := ioutil.ReadDir(pkg.YummyPath)
		if err != nil {
			myErr.Logging(err)
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

func (monitor *Monitor) passedCheck(currentfiles, newFiles []string) {
	//最新のリストからファイルが削除されていれば渡しずみ判断
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

	return
}
