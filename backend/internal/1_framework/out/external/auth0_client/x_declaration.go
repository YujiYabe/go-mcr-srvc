package auth0_client

import (
	externalGateway "backend/internal/2_adapter/gateway/external"
)

// Auth0Client ...
type Auth0Client struct{}

// NewToAuth0 ...
func NewToAuth0() (
	toAuth0 externalGateway.ToAuth0,
) {
	toAuth0 = new(Auth0Client)
	return
}
