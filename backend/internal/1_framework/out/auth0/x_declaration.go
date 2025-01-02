package auth0

import (
	"backend/internal/2_adapter/gateway"
)

// Auth0 ...
type Auth0 struct{}

// NewToAuth0 ...
func NewToAuth0() (
	toAuth0 gateway.ToAuth0,
) {
	toAuth0 = new(Auth0)
	return
}
