package auth0_client

import (
	"backend/internal/2_adapter/gateway"
)

// Auth0Client ...
type Auth0Client struct{}

// NewToAuth0 ...
func NewToAuth0() (
	toAuth0 gateway.ToAuth0,
) {
	toAuth0 = new(Auth0Client)
	return
}
