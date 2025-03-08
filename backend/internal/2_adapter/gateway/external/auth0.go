package external_gateway

import (
	"context"

	domainObject "backend/internal/4_domain/domain_object"
	groupObject "backend/internal/4_domain/group_object"
)

// FetchAccessToken ...
func (receiver *GatewayExternal) FetchAccessToken(
	ctx context.Context,
	credential groupObject.Credential,
) (
	accessToken domainObject.AccessToken,
) {
	accessToken = receiver.ToAuth0.FetchAccessToken(
		ctx,
		credential,
	)
	return
}
