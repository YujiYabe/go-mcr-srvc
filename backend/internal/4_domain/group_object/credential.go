package group_object

import valueObject "backend/internal/4_domain/value_object"

type Credential struct {
	Err          error
	ClientID     valueObject.ClientID
	ClientSecret valueObject.ClientSecret
}

type NewCredentialArgs struct {
	ClientID     *string
	ClientSecret *string
}

func (receiver *Credential) GetError() error {
	return receiver.Err
}

func (receiver *Credential) SetError(
	err error,
) {
	if receiver.Err == nil {
		receiver.Err = err
	}
}

func NewCredential(
	args *NewCredentialArgs,
) (
	accessToken *Credential,
) {
	accessToken = &Credential{}

	accessToken.ClientID = valueObject.NewClientID(args.ClientID)
	if accessToken.ClientID.Err != nil {
		accessToken.SetError(accessToken.ClientID.Err)
		return
	}

	accessToken.ClientSecret = valueObject.NewClientSecret(args.ClientSecret)
	if accessToken.ClientSecret.Err != nil {
		accessToken.SetError(accessToken.ClientSecret.Err)
		return
	}

	return
}
