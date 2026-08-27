package usecase

import (
	"context"
	"fmt"

	groupObject "backend/internal/4_domain/group_object"
	typeObject "backend/internal/4_domain/type_object"
)

func (receiver *useCase) FetchAccessToken(
	ctx context.Context,
	credential groupObject.Credential,
) (
	accessToken typeObject.AccessToken,
	err error,
) {
	if err = ensureContextReady(ctx, "FetchAccessToken"); err != nil {
		return
	}
	if err = credential.EnsureReadyToAuthenticate(); err != nil {
		err = fmt.Errorf("FetchAccessToken: %w", err)
		return
	}
	accessToken, err = receiver.ToGatewayExternal.FetchAccessToken(
		ctx,
		credential,
	)
	if err != nil {
		err = fmt.Errorf("FetchAccessToken: %w", err)
	}
	return
}
