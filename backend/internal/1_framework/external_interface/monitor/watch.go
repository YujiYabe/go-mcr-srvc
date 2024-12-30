package monitor

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"backend/pkg"
)

func (receiver *Monitor) RemoveYummy() {
	ctx := context.Background()
	yummyFiles, err := os.ReadDir(pkg.YummyPath)
	if err != nil {
		pkg.Logging(ctx, err)
	}

	for _, yummyFile := range yummyFiles {
		if yummyFile.Name() == "readme.md" {
			continue
		}
		err := os.Remove(filepath.Join(pkg.YummyPath, yummyFile.Name()))
		if err != nil {
			pkg.Logging(ctx, err)
		}
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
