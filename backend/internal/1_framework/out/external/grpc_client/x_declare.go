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
	toGRPC = gatewayExternal.ToGRPC(nil)
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

		conn, returnedErr := grpc.NewClient(
			address,
			grpc.WithTransportCredentials(
				insecure.NewCredentials(),
			),
		)
		if returnedErr == nil {
			return &GRPCClient{
				Conn: conn,
			}, nil
		}

		lastErr = returnedErr
		logger.Logging(ctx, returnedErr)

		if attempt == count {
			break
		}

		select {
		case <-ctx.Done():
			gRPCClient, err = nil, ctx.Err()
			return
		case <-time.After(retryBackoff(attempt)):
		}
	}

	gRPCClient, err = nil, fmt.Errorf("retry count over: %w", lastErr)
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
