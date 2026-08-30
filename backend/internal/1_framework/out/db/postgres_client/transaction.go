package postgres_client

import (
	"context"

	"gorm.io/gorm"
)

type txContextKey struct{}

func (receiver *PostgresClient) RunInTransaction(
	ctx context.Context,
	fn func(context.Context) error,
) (
	err error,
) {
	return receiver.Conn.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(context.WithValue(ctx, txContextKey{}, tx))
	})
}

func (receiver *PostgresClient) conn(
	ctx context.Context,
) (
	dB *gorm.DB,
) {
	if tx, ok := ctx.Value(txContextKey{}).(*gorm.DB); ok {
		return tx.WithContext(ctx)
	}
	return receiver.Conn.WithContext(ctx)
}
