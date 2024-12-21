package monitor

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"time"

	"backend/pkg"
)

func (receiver *Monitor) RemoveYummy() {
	yummyFiles, err := os.ReadDir(pkg.YummyPath)
	if err != nil {
		myErr.Logging(err)
	}

	for _, yummyFile := range yummyFiles {
		if yummyFile.Name() == "readme.md" {
			continue
		}
		err := os.Remove(filepath.Join(pkg.YummyPath, yummyFile.Name()))
		if err != nil {
			myErr.Logging(err)
		}
	}

}

func (receiver *Monitor) Watching() {
	var currentFiles []string

	for {
		files, err := os.ReadDir(pkg.YummyPath)
		if err != nil {
			myErr.Logging(err)
		}

		var newFiles []string
		for _, file := range files {
			newFiles = append(newFiles, file.Name())
		}

		receiver.passedCheck(currentFiles, newFiles)

		currentFiles = newFiles

		time.Sleep(1 * time.Second)
	}
}

func (receiver *Monitor) passedCheck(currentFiles, newFiles []string) {
	//最新のリストからファイルが削除されていれば渡しずみ判断
	for _, currentFile := range currentFiles {
		isExist := false
		for _, newFile := range newFiles {
			if currentFile == newFile {
				isExist = true
				continue
			}
		}

		if !isExist {
			ctx := context.Background()
			receiver.UpdateOrders(ctx, strings.TrimRight(currentFile, ".json"), "pass")
		}
	}

}
