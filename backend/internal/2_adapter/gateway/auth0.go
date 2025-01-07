package gateway

import (
	"context"

	groupObject "backend/internal/4_domain/group_object"
	valueObject "backend/internal/4_domain/value_object"
)

// FetchAccessToken ...
func (receiver *Gateway) FetchAccessToken(
	ctx context.Context,
	credential groupObject.Credential,
) (
	accessToken valueObject.AccessToken,
) {
	accessToken = receiver.ToAuth0.FetchAccessToken(
		ctx,
		credential,
	)
	return
}
