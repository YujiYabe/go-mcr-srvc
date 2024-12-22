package mysql

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"backend/internal/2_adapter/gateway"
	"backend/pkg"
)

var isEnable = false

type (
	// MySQL ...
	MySQL struct {
		Conn *gorm.DB
	}

	// Vegetable ...
	Vegetable struct {
		ID    int
		Name  string
		Stock int
	}
)

// NewToMySQL ...
func NewToMySQL() gateway.ToMySQL {
	ctx := context.Background()
	conn, err := open(30)
	if err != nil {
		pkg.Logging(ctx, err)
		panic(err)
	}

	s := new(MySQL)
	s.Conn = conn
	return s
}

func open(count uint) (
	client *gorm.DB,
	err error,

) {
	if !isEnable {
		return client, nil
	}

	ctx := context.Background()
	client, err = gorm.Open(mysql.Open(pkg.MySQLDSN), &gorm.Config{})
	if err != nil {
		if count == 0 {
			pkg.Logging(ctx, err)
			return nil, fmt.Errorf("Retry count over")
		}
		time.Sleep(time.Second)
		// カウントダウンさせるようにする
		count--
		return open(count)
	}

	return client, nil
}

// UpdatePatties ...
func (receiver *MySQL) UpdatePatties(ctx context.Context, items map[string]int) error {
	for item, num := range items {
		res := receiver.Conn.
			Table("patties").
			Where("name IN (?)", item).
			UpdateColumn("stock", gorm.Expr("stock - ?", num))

		if res.Error != nil {
			pkg.Logging(ctx, res.Error)
			return res.Error
		}

		// 作業時間を擬似的に再現
		time.Sleep(1 * time.Second)
	}
	return nil
}
