package gateway

import (
	"context"

	structObject "backend/internal/4_domain/struct_object"
	valueObject "backend/internal/4_domain/value_object"
)

// FetchAccessToken ...
func (receiver *Gateway) FetchAccessToken(
	ctx context.Context,
	credential structObject.Credential,
) (
	accessToken valueObject.AccessToken,
	err error,
) {
	accessToken, err = receiver.ToAuth0.FetchAccessToken(
		ctx,
		credential,
	)
	return
}
