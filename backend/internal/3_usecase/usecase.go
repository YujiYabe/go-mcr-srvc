package usecase

import (
	"context"
	"fmt"
)

func ensureContextReady(
	ctx context.Context,
	usecaseName string,
) (
	err error,
) {
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("%s: context is not ready: %w", usecaseName, err)
	}

	return nil
}
