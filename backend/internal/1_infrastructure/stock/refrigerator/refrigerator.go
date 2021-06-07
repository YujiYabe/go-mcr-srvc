package refrigerator

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"backend/internal/2_adapter/service"
	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("infrastructure", "refrigerator")
}

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
		myErr.Logging(err)
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
			myErr.Logging(err)
			return nil, fmt.Errorf("Retry count over")
		}
		time.Sleep(time.Second)
		count--
		return open(count)
	}

	return db, nil
}

// GetVegetables ...
func (s *Refrigerator) GetVegetables(ctx context.Context, items map[string]int) error {
	for item, num := range items {
		res := s.Conn.
			Table("vegetables").
			Where("name IN (?)", item).
			UpdateColumn("stock", gorm.Expr("stock - ?", num))

		if res.Error != nil {
			myErr.Logging(res.Error)
			return res.Error
		}

		time.Sleep(2 * time.Second)
	}

	return nil
}

// GetIngredients ...
func (s *Refrigerator) GetIngredients(ctx context.Context, items map[string]int) error {
	for item, num := range items {
		res := s.Conn.
			Table("ingredients").
			Where("name IN (?)", item).
			UpdateColumn("stock", gorm.Expr("stock - ?", num))

		if res.Error != nil {
			myErr.Logging(res.Error)
			return res.Error
		}
	}

	return nil
}
