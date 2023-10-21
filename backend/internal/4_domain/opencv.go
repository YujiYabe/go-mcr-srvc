package domain

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/anikhasibul/queue"
	"gocv.io/x/gocv"
)

func CodeDetector(number int, file *multipart.FileHeader) ([]int, error) {
	janCodes := []int{}
	src, err := file.Open()
	fileName := fmt.Sprintf("%d-%d-%s", number, time.Now().Unix(), file.Filename)

	if err != nil {
		return janCodes, err
	}
	defer src.Close()

	img, err := loadImage(src, fileName)
	if err != nil {
		return janCodes, err
	}
	defer img.Close()

	janCodes, err = detectQRCode(img)
	if err != nil {
		return janCodes, err
	}

	return janCodes, nil
}

func loadImage(src multipart.File, fileName string) (gocv.Mat, error) {
	var img gocv.Mat

	// 画像の読み込み方法を決定
	// true  画像をメモリで処理してファイルとして保存しない
	// false 画像ファイルとして保存してから処理する→デバッグ用
	if true {
		fileBytes, err := io.ReadAll(src)
		if err != nil {
			return img, err
		}

		img, err = gocv.IMDecode(fileBytes, gocv.IMReadGrayScale)
		if err != nil {
			return img, err
		}
	} else {
		tempDir := "uploads"
		ensureDirExists(tempDir)

		dstPath := filepath.Join(tempDir, fileName)
		dst, err := os.Create(dstPath)
		if err != nil {
			return img, err
		}
		defer dst.Close()

		// 画像ファイルのコピー
		if _, err = io.Copy(dst, src); err != nil {
			return img, err
		}

		img = gocv.IMRead(dstPath, gocv.IMReadGrayScale)
		if img.Empty() {
			return img, err
		}
	}

	return img, nil
}

func detectQRCode(img gocv.Mat) ([]int, error) {
	janCodes := []int{}

	// QRコードの検出器を初期化
	qcd := gocv.NewQRCodeDetector()
	defer qcd.Close()

	// 処理の開始時刻を記録
	startTime := time.Now()
	fmt.Printf("start time: %v\n", startTime)

	// 並行処理のためのキューの初期化
	capacity := 10
	q := queue.New(capacity) // 擬似的に同時進行できるキャパシティを設定
	defer q.Close()

	for i := 0; i < capacity; i++ {
		q.Add()
		go func(c int) {
			defer q.Done()
			tmpCodes := PickOutJANCodes(qcd, img)
			if len(janCodes) <= len(tmpCodes) {
				janCodes = tmpCodes
			}
		}(i)
	}
	q.Wait()
	fmt.Printf("process time: %s\n", time.Since(startTime))

	return janCodes, nil
}

// ensureDirExists はディレクトリが存在することを確認し、存在しない場合は作成します。
func ensureDirExists(dirName string) {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		os.Mkdir(dirName, 0755)
	}
}
