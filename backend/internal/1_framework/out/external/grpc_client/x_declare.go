package grpc_client

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	gatewayExternal "backend/internal/2_adapter/gateway/external"
	"backend/internal/logger"
)

// GRPCClient ...
type (
	GRPCClient struct {
		Conn *grpc.ClientConn
	}
)

// NewToGRPC ...
func NewToGRPC(
	ctx context.Context,
	address string,
) (
	toGRPC gatewayExternal.ToGRPC,
	err error,
) {
	conn, err := open(ctx, address, 30)
	if err != nil {
		return
	}

	toGRPC = conn

	return
}

func open(
	ctx context.Context,
	address string,
	count uint,
) (
	gRPCClient *GRPCClient,
	err error,
) {
	var lastErr error
	for attempt := uint(0); attempt <= count; attempt++ {
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		conn, err := grpc.NewClient(
			address,
			grpc.WithTransportCredentials(
				insecure.NewCredentials(),
			),
		)
		if err == nil {
			return &GRPCClient{
				Conn: conn,
			}, nil
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
