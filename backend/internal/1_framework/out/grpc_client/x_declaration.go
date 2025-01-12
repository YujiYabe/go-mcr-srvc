package grpc_client

import (
	"backend/internal/2_adapter/gateway"
	"backend/pkg"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
		pkg.Logging(ctx, err)
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
			pkg.Logging(ctx, err)
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
