package group_object

import (
	"context"

	typeObject "backend/internal/4_domain/type_object"
	"backend/internal/logger"
)

type Credential struct {
	err          error
	ClientID     typeObject.ClientID
	ClientSecret typeObject.ClientSecret
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

	accessToken.ClientID = typeObject.NewClientID(
		ctx,
		args.ClientID,
	)
	if accessToken.ClientID.GetError() != nil {
		logger.Logging(ctx, accessToken.ClientID.GetError())
		accessToken.SetError(ctx, accessToken.ClientID.GetError())
		return
	}

	accessToken.ClientSecret = typeObject.NewClientSecret(ctx, args.ClientSecret)
	if accessToken.ClientSecret.GetError() != nil {
		logger.Logging(ctx, accessToken.ClientSecret.GetError())
		accessToken.SetError(ctx, accessToken.ClientSecret.GetError())
		return
	}

	return
}
