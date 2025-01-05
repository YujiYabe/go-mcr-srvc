package struct_object

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
) *Credential {
	if receiver.Err == nil {
		receiver.Err = err
	}

	return receiver
}

func NewCredential(
	args *NewCredentialArgs,
) (
	accessToken *Credential,
) {
	var err error
	accessToken = &Credential{}

	accessToken.ClientID, err = valueObject.NewClientID(args.ClientID)
	if err != nil {
		accessToken.SetError(err)
		return
	}

	accessToken.ClientSecret, err = valueObject.NewClientSecret(args.ClientSecret)
	if err != nil {
		accessToken.SetError(err)
		return
	}

	return
}
