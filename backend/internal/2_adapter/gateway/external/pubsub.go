package external_gateway

import (
	"context"
)

// ToPubSub ...
func (receiver *GatewayExternal) PublishTestTopic(
	ctx context.Context,
) (
	err error,
) {
	return receiver.ToPubSub.PublishTestTopic(
		ctx,
	)
}
