package postgres

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"backend/internal/2_adapter/gateway"
	"backend/pkg"
)

var (
	myErr *pkg.MyErr
)

func init() {
	myErr = pkg.NewMyErr("framework_driver", "postgres")
}

type (
	// Postgres ...
	Postgres struct {
		Conn *gorm.DB
	}

	// Vegetable ...
	Vegetable struct {
		ID    int
		Name  string
		Stock int
	}
)

// NewToPostgres ...
func NewToPostgres() gateway.ToPostgres {
	conn, err := open(30)
	if err != nil {
		myErr.Logging(err)
		panic(err)
	}

	s := new(Postgres)
	s.Conn = conn
	return s
}

func open(count uint) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(pkg.PostgresDSN), &gorm.Config{})

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

// UpdateVegetables ...
func (receiver *Postgres) UpdateVegetables(ctx context.Context, items map[string]int) error {
	for item, num := range items {
		res := receiver.Conn.
			Table("vegetables").
			Where("name IN (?)", item).
			UpdateColumn("stock", gorm.Expr("stock - ?", num))

		if res.Error != nil {
			myErr.Logging(res.Error)
			return res.Error
		}

		// 作業時間を擬似的に再現
		time.Sleep(1 * time.Second)
	}

	return nil
}

// UpdateIngredients ...
func (receiver *Postgres) UpdateIngredients(ctx context.Context, items map[string]int) error {
	for item, num := range items {
		res := receiver.Conn.
			Table("ingredients").
			Where("name IN (?)", item).
			UpdateColumn("stock", gorm.Expr("stock - ?", num))

		if res.Error != nil {
			myErr.Logging(res.Error)
			return res.Error
		}

		// 作業時間を擬似的に再現
		time.Sleep(1 * time.Second)
	}

	return nil
}
