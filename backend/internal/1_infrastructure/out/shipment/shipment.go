package shipment

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"app/internal/2_adapter/service"
	"app/internal/4_domain/domain"
)

type Shipment struct{}

// NewToShipment ...
func NewToShipment() service.ToShipment {
	s := new(Shipment)
	return s
}

// HandOver ...
func (s *Shipment) HandOver(ctx context.Context, order *domain.Order) error {
	return nil
}

// Logging ...
func (s *Shipment) Logging(ctx context.Context, order *domain.Order) error {
	currentPath, _ := os.Getwd()
	LogPath := filepath.Join(currentPath, "storage", "log")

	fileName := time.Now().Format("2006-01-02") + ".log"
	LogName := filepath.Join(LogPath, fileName)

	// ファイルが存在しなければ作成
	_, err := os.Stat(LogName)
	if err != nil {
		_, err := os.Create(LogName)
		if err != nil {
			log.Fatal(err)
		}
	}

	// エラー情報作成
	var rowData []string
	rowData = append(rowData, order.OrderNumber)                             // レイヤ名
	rowData = append(rowData, order.OrderTime.Format("2006-01-02_15:04:05")) // オーダー時間
	rowData = append(rowData, time.Now().Format("2006-01-02_15:04:05"))      // 引き渡し時間
	rowData = append(rowData, order.OrderType)                               // オーダータイプ
	row := strings.Join(rowData, "\t")                                       // タブ区切り

	// ファイル書き込み
	f, err := os.OpenFile(filepath.Clean(LogName), os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Fprintln(f, row)

	return nil
}
