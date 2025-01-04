package gateway

import (
	"context"
)

// ViaGRPC ...
func (receiver *Gateway) ViaGRPC(
	ctx context.Context,
) (
	err error,
) {
	err = receiver.ToGRPC.ViaGRPC(
		ctx,
	)

	return
}
