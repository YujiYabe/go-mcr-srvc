package auth0_client

import (
	gatewayExternal "backend/internal/2_adapter/gateway/external"
)

// Auth0Client ...
type Auth0Client struct{}

// NewToAuth0 ...
func NewToAuth0() (
	toAuth0 gatewayExternal.ToAuth0,
) {
	toAuth0 = new(Auth0Client)
	return
}
