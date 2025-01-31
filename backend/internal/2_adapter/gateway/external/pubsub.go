package external_gateway

import (
	"context"
)

// ToPubSub ...
func (receiver *GatewayExternal) PublishTestTopic(
	ctx context.Context,

) {
	receiver.ToPubSub.PublishTestTopic(
		ctx,
	)
}
