package external_gateway

import (
	"context"
)

// ToPubSub ...
func (receiver *GatewayExternal) PublishTestTopic(
	ctx context.Context,
) error {
	return receiver.ToPubSub.PublishTestTopic(
		ctx,
	)
}
