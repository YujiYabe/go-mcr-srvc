package auth0

import (
	"context"

	"backend/internal/4_domain/struct_object"
	"backend/internal/4_domain/value_object"
)

// ...
// GetAccessToken ...
func (receiver *Auth0) GetAccessToken(
	ctx context.Context,
	credential struct_object.Credential,
) (
	accessToken value_object.AccessToken,
	err error,
) {

	return
}
