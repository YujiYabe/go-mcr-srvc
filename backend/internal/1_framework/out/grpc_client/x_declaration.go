package grpc_client

import (
	"backend/internal/2_adapter/gateway"
)

// GRPCClient ...
type GRPCClient struct{}

// NewToGRPC ...
func NewToGRPC() (
	toGRPC gateway.ToGRPC,
) {
	toGRPC = new(GRPCClient)
	return
}
