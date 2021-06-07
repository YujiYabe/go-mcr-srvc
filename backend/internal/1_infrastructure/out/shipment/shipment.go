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
	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("infrastructure", "shipment")
}

type Shipment struct{}

// NewToShipment ...
func NewToShipment() service.ToShipment {
	s := new(Shipment)
	return s
}

// HandOver ...
func (s *Shipment) HandOver(ctx context.Context, order *domain.Order) error {
	fileName := order.OrderInfo.OrderNumber + ".json"
	yummyFilePath := filepath.Join(pkg.YummyPath, fileName)

	product, err := json.MarshalIndent(order.Product, "", "    ")
	if err != nil {
		myErr.Logging(err)
		return err
	}

	err = ioutil.WriteFile(yummyFilePath, product, 0777)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	return nil
}

// Logging ...
func (s *Shipment) Logging(ctx context.Context, order *domain.Order) error {
	fileName := time.Now().Format("2006-01-02") + ".log"
	LogName := filepath.Join(pkg.LogPath, fileName)

	// ファイルが存在しなければ作成
	_, err := os.Stat(LogName)
	if err != nil {
		_, err := os.Create(LogName)
		if err != nil {
			myErr.Logging(err)
			return err
		}
	}

	// エラー以外の情報をjson化
	product, err := json.Marshal(&order.Product)
	if err != nil {
		myErr.Logging(err)
		return err
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
		myErr.Logging(err)
		return err
	}

	defer func() {
		err := f.Close()
		if err != nil {
			myErr.Logging(err)
			log.Fatal(err)
		}
	}()

	fmt.Fprintln(f, row)

	return nil
}
