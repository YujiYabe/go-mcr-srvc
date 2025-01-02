package gateway

import (
	"context"

	"backend/internal/4_domain/struct_object"
	"backend/internal/4_domain/value_object"
)

// FetchAccessToken ...
func (receiver *Gateway) FetchAccessToken(
	ctx context.Context,
	credential struct_object.Credential,
) (
	accessToken value_object.AccessToken,
	err error,
) {
	//  receiver.ToPostgres.GetPersonByCondition(
	// 	ctx,
	// 	reqPerson,
	// )

	return
}
