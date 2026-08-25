package group_object

import typeObject "backend/internal/4_domain/type_object"

type Credential struct {
	clientID     typeObject.ClientID
	clientSecret typeObject.ClientSecret
}

type NewCredentialArgs struct {
	ClientID     *string
	ClientSecret *string
}

func (receiver Credential) ClientID() typeObject.ClientID {
	return receiver.clientID
}

func (receiver Credential) ClientSecret() typeObject.ClientSecret {
	return receiver.clientSecret
}

func NewCredential(
	args *NewCredentialArgs,
) (
	credential *Credential,
	err error,
) {
	credential = &Credential{}

	credential.clientID, err = typeObject.NewClientID(
		args.ClientID,
	)
	if err != nil {
		return nil, err
	}

	credential.clientSecret, err = typeObject.NewClientSecret(args.ClientSecret)
	if err != nil {
		return nil, err
	}

	return
}
