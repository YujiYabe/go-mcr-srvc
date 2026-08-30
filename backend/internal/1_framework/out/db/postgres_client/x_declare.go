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
		return nil, err
	}

	postgresClient := new(PostgresClient)
	postgresClient.Conn = conn
	return postgresClient, nil
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

		db, err := gorm.Open(
			postgres.Open(dsn),
			&gorm.Config{},
		)
		if err == nil {
			return db, nil
		}

		lastErr = err
		logger.Logging(ctx, err)

		if attempt == count {
			break
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryBackoff(attempt)):
		}
	}

	return nil, fmt.Errorf("retry count over: %w", lastErr)
}

func retryBackoff(
	attempt uint,
) (
	duration time.Duration,
) {
	backoff := time.Duration(attempt+1) * time.Second
	if backoff > 5*time.Second {
		return 5 * time.Second
	}
	return backoff
}
