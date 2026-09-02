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
	err = receiver.ToPubSub.PublishTestTopic(
		ctx,
	)
	return
}
