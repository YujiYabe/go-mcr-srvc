package usecase

import (
	"context"
	"fmt"
)

func (receiver *useCase) PublishTestTopic(
	ctx context.Context,
) (
	err error,
) {
	if returnedErr := ensureContextReady(ctx, "PublishTestTopic"); returnedErr != nil {
		err = returnedErr
		return
	}
	if err := receiver.ToGatewayExternal.PublishTestTopic(ctx); err != nil {
		return fmt.Errorf("PublishTestTopic: %w", err)
	}

	err = nil
	return
}
