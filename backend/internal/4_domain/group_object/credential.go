package group_object

import (
	"context"

	valueObject "backend/internal/4_domain/value_object"
	"backend/pkg"
)

type Credential struct {
	err          error
	ClientID     valueObject.ClientID
	ClientSecret valueObject.ClientSecret
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

	accessToken.ClientID = valueObject.NewClientID(
		ctx,
		args.ClientID,
	)
	if accessToken.ClientID.GetError() != nil {
		pkg.Logging(ctx, accessToken.ClientID.GetError())
		accessToken.SetError(ctx, accessToken.ClientID.GetError())
		return
	}

	accessToken.ClientSecret = valueObject.NewClientSecret(ctx, args.ClientSecret)
	if accessToken.ClientSecret.GetError() != nil {
		pkg.Logging(ctx, accessToken.ClientSecret.GetError())
		accessToken.SetError(ctx, accessToken.ClientSecret.GetError())
		return
	}

	return
}
