package grpc_client

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"backend/internal/2_adapter/gateway"
	logger "backend/internal/logger"
)

// GRPCClient ...
type (
	GRPCClient struct {
		Conn *grpc.ClientConn
	}
)

// NewToGRPC ...
func NewToGRPC() (
	toGRPC gateway.ToGRPC,
) {
	ctx := context.Background()
	conn, err := open(ctx, 30)
	if err != nil {
		logger.Logging(ctx, err)
		panic(err)
	}

	toGRPC = conn

	return
}

func open(
	ctx context.Context,
	count uint,
) (*GRPCClient, error) {
	conn, err := grpc.NewClient(
		"backend:3456",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)

	if err != nil {
		if count == 0 {
			logger.Logging(ctx, err)
			return nil, fmt.Errorf("retry count over")
		}
		time.Sleep(time.Second)
		count--
		return open(ctx, count)
	}

	return &GRPCClient{
		Conn: conn,
	}, nil
}
