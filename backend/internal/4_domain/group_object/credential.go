package group_object

import (
	"context"

	domainObject "backend/internal/4_domain/domain_object"
	"backend/internal/logger"
)

type Credential struct {
	err          error
	ClientID     domainObject.ClientID
	ClientSecret domainObject.ClientSecret
}

type NewCredentialArgs struct {
	ClientID     *string
	ClientSecret *string
}

func (receiver *Credential) GetError() error {
	return receiver.err
}

func (receiver *Credential) SetError(
	ctx context.Context, err error,
) {
	if receiver.err == nil {
		receiver.err = err
	}
}

func NewCredential(
	ctx context.Context,
	args *NewCredentialArgs,
) (
	accessToken *Credential,
) {
	accessToken = &Credential{}

	accessToken.ClientID = domainObject.NewClientID(
		ctx,
		args.ClientID,
	)
	if accessToken.ClientID.GetError() != nil {
		logger.Logging(ctx, accessToken.ClientID.GetError())
		accessToken.SetError(ctx, accessToken.ClientID.GetError())
		return
	}

	accessToken.ClientSecret = domainObject.NewClientSecret(ctx, args.ClientSecret)
	if accessToken.ClientSecret.GetError() != nil {
		logger.Logging(ctx, accessToken.ClientSecret.GetError())
		accessToken.SetError(ctx, accessToken.ClientSecret.GetError())
		return
	}

	return
}
