package postgres

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"backend/internal/2_adapter/gateway"
	"backend/internal/4_domain/struct_object"
	"backend/pkg"
)

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
	ctx := context.Background()
	conn, err := open(30)
	if err != nil {
		pkg.Logging(ctx, err)
		panic(err)
	}

	postgres := new(Postgres)
	postgres.Conn = conn
	return postgres
}

func open(count uint) (*gorm.DB, error) {
	ctx := context.Background()
	db, err := gorm.Open(postgres.Open(pkg.PostgresDSN), &gorm.Config{})

	if err != nil {
		if count == 0 {
			pkg.Logging(ctx, err)
			return nil, fmt.Errorf("retry count over")
		}
		time.Sleep(time.Second)
		count--
		return open(count)
	}

	return db, nil
}

// GetPersonList ...
func (receiver *Postgres) GetPersonList(
	ctx context.Context,
) (
	personList struct_object.PersonList,
	err error,
) {
	personList = struct_object.PersonList{}

	res := receiver.Conn.
		Table("persons").
		Find(personList)

	if res.Error != nil {
		pkg.Logging(ctx, res.Error)
		return personList, res.Error
	}

	return personList, nil
}
