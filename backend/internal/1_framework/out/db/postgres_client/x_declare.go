package postgres_client

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	gatewayDB "backend/internal/2_adapter/gateway/db"
	"backend/internal/env"
	"backend/internal/logger"
)

type (
	// PostgresClient ...
	PostgresClient struct {
		Conn *gorm.DB
	}
)

// NewToPostgres ...
func NewToPostgres() gatewayDB.ToPostgres {
	ctx := context.Background()
	conn, err := open(ctx, 30)
	if err != nil {
		logger.Logging(ctx, err)
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
	db, err := gorm.Open(
		postgres.Open(env.DatabaseConfig.DSN),
		&gorm.Config{},
	)

	if err != nil {
		if count == 0 {
			logger.Logging(ctx, err)
			return nil, fmt.Errorf(
				"retry count over")
		}
		time.Sleep(time.Second)
		count--
		return open(ctx, count)
	}

	return db, nil
}
