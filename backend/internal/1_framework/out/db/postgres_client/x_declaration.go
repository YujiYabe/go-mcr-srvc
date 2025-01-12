package postgres_client

import (
	"backend/pkg"
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"backend/internal/2_adapter/gateway"
)

type (
	// PostgresClient ...
	PostgresClient struct {
		Conn *gorm.DB
	}
)

// NewToPostgres ...
func NewToPostgres() gateway.ToPostgres {
	ctx := context.Background()
	conn, err := open(ctx, 30)
	if err != nil {
		pkg.Logging(ctx, err)
		panic(err)
	}

	postgresClient := new(PostgresClient)
	postgresClient.Conn = conn
	return postgresClient
}

func open(
	ctx context.Context,
	count uint,
) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(pkg.PostgresDSN), &gorm.Config{})

	if err != nil {
		if count == 0 {
			pkg.Logging(ctx, err)
			return nil, fmt.Errorf(
				"retry count over")
		}
		time.Sleep(time.Second)
		count--
		return open(ctx, count)
	}

	return db, nil
}
