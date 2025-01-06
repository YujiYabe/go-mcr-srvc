package group_object

import valueObject "backend/internal/4_domain/value_object"

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
	err error,
) {
	if receiver.err == nil {
		receiver.err = err
	}
}

func NewCredential(
	args *NewCredentialArgs,
) (
	accessToken *Credential,
) {
	accessToken = &Credential{}

	accessToken.ClientID = valueObject.NewClientID(args.ClientID)
	if accessToken.ClientID.GetError() != nil {
		accessToken.SetError(accessToken.ClientID.GetError())
		return
	}

	accessToken.ClientSecret = valueObject.NewClientSecret(args.ClientSecret)
	if accessToken.ClientSecret.GetError() != nil {
		accessToken.SetError(accessToken.ClientSecret.GetError())
		return
	}

	return
}
