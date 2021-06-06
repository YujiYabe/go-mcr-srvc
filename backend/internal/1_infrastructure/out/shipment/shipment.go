package shipment

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"backend/internal/2_adapter/service"
	"backend/internal/4_domain/domain"
)

type Shipment struct{}

// NewToShipment ...
func NewToShipment() service.ToShipment {
	s := new(Shipment)
	return s
}

// HandOver ...
func (s *Shipment) HandOver(ctx context.Context, order *domain.Order) error {
	currentPath, _ := os.Getwd()
	fileName := order.OrderInfo.OrderNumber + ".json"

	yummyPath := filepath.Join(currentPath, "yummy", fileName)

	product, err := json.MarshalIndent(order.Product, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(yummyPath, product, 0777)
	if err != nil {
		return err
	}

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

	// エラー以外の情報をjson化
	product, err := json.Marshal(&order.Product)
	if err != nil {
		log.Fatal(err)
	}

	// エラー情報作成
	var rowData []string
	rowData = append(rowData, order.OrderInfo.OrderNumber)                             // レイヤ名
	rowData = append(rowData, order.OrderInfo.OrderTime.Format("2006-01-02_15:04:05")) // オーダー時間
	rowData = append(rowData, time.Now().Format("2006-01-02_15:04:05"))                // 引き渡し時間
	rowData = append(rowData, order.OrderInfo.OrderType)                               // オーダータイプ
	rowData = append(rowData, string(product))                                         // 商品
	row := strings.Join(rowData, "\t")                                                 // タブ区切り

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
