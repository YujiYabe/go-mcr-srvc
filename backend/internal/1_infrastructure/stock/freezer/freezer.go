package freezer

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"backend/internal/2_adapter/service"
)

type (
	// Freezer ...
	Freezer struct {
		Conn *gorm.DB
	}

	// Vegetable ...
	Vegetable struct {
		ID    int
		Name  string
		Stock int
	}
)

// NewToFreezer ...
func NewToFreezer() service.ToFreezer {
	conn, err := open(30)
	if err != nil {
		panic(err)
	}

	s := new(Freezer)
	s.Conn = conn
	return s
}

func open(count uint) (*gorm.DB, error) {
	dsn := "user:user@tcp(mysql)/app?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		if count == 0 {
			return nil, fmt.Errorf("Retry count over")
		}
		time.Sleep(time.Second)
		// カウントダウンさせるようにする
		count--
		return open(count)
	}

	return db, nil
}

// Dummy ...
func (s *Freezer) Dummy(ctx context.Context) (string, error) {
	return "dummy ok", nil
}

// GetPatties ...
func (s *Freezer) GetPatties(ctx context.Context, items map[string]int) error {
	for item, num := range items {
		res := s.Conn.
			Table("patties").
			Where("name IN (?)", item).
			UpdateColumn("stock", gorm.Expr("stock - ?", num))

		if res.Error != nil {
			return res.Error
		}
	}

	return nil
}
