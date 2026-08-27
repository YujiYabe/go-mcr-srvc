package usecase

import (
	"context"
	"fmt"
)

func (receiver *useCase) PublishTestTopic(
	ctx context.Context,
) error {
	if err := ensureContextReady(ctx, "PublishTestTopic"); err != nil {
		return err
	}
	if err := receiver.ToGatewayExternal.PublishTestTopic(ctx); err != nil {
		return fmt.Errorf("PublishTestTopic: %w", err)
	}

	return nil
}
