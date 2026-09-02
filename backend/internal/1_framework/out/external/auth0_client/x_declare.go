package auth0_client

import (
	gatewayExternal "backend/internal/2_adapter/gateway/external"
)

// Auth0Client ...
type Auth0Client struct {
	tokenURL  string
	audience  string
	grantType string
}

// NewToAuth0 ...
func NewToAuth0(
	tokenURL string,
	audience string,
	grantType string,
) (
	toAuth0 gatewayExternal.ToAuth0,
) {
	toAuth0 = &Auth0Client{
		tokenURL:  tokenURL,
		audience:  audience,
		grantType: grantType,
	}

	return
}
