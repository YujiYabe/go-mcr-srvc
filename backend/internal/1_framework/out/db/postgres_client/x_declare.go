package postgres_client

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	gatewayDB "backend/internal/2_adapter/gateway/db"
	"backend/internal/logger"
)

type (
	// PostgresClient ...
	PostgresClient struct {
		Conn *gorm.DB
	}
)

// NewToPostgres ...
func NewToPostgres(
	ctx context.Context,
	dsn string,
) (
	toPostgres gatewayDB.ToPostgres,
	err error,
) {
	conn, err := open(ctx, dsn, 30)
	if err != nil {
		toPostgres = nil

		return
	}

	postgresClient := new(PostgresClient)
	postgresClient.Conn = conn
	toPostgres, err = postgresClient, nil

	return
}

func open(
	ctx context.Context,
	dsn string,
	count uint,
) (
	dB *gorm.DB,
	err error,
) {
	var lastErr error
	for attempt := uint(0); attempt <= count; attempt++ {
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		db, returnedErr := gorm.Open(
			postgres.Open(dsn),
			&gorm.Config{},
		)
		if returnedErr == nil {
			return db, nil
		}

		lastErr = returnedErr
		logger.Logging(ctx, returnedErr)

		if attempt == count {
			break
		}

		select {
		case <-ctx.Done():
			dB, err = nil, ctx.Err()
			return
		case <-time.After(retryBackoff(attempt)):
		}
	}

	dB, err = nil, fmt.Errorf("retry count over: %w", lastErr)
	return
}

func retryBackoff(
	attempt uint,
) (
	duration time.Duration,
) {
	if attempt >= 4 {
		duration = 5 * time.Second

		return
	}

	duration = time.Duration(attempt+1) * time.Second

	return
}
