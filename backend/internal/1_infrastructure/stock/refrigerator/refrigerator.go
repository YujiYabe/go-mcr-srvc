package refrigerator

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"app/internal/2_adapter/service"
)

type (
	// Refrigerator ...
	Refrigerator struct {
		Conn *gorm.DB
	}

	// Vegetable ...
	Vegetable struct {
		ID    int
		Name  string
		Stock int
	}
)

// NewToRefrigerator ...
func NewToRefrigerator() service.ToRefrigerator {
	conn, err := open(30)
	if err != nil {
		panic(err)
	}

	s := new(Refrigerator)
	s.Conn = conn
	return s
}

func open(count uint) (*gorm.DB, error) {
	dsn := "host=postgres user=user password=user dbname=app port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		if count == 0 {
			return nil, fmt.Errorf("Retry count over")
		}
		time.Sleep(time.Second)
		count--
		return open(count)
	}

	return db, nil
}

// Dummy ...
func (s *Refrigerator) Dummy(ctx context.Context) (string, error) {
	return "dummy ok", nil
}

// GetVegetables ...
func (s *Refrigerator) GetVegetables(ctx context.Context, items map[string]int) error {
	for item, num := range items {
		res := s.Conn.
			Table("vegetables").
			Where("name IN (?)", item).
			UpdateColumn("stock", gorm.Expr("stock - ?", num))

		if res.Error != nil {
			return res.Error
		}
	}

	return nil
}
