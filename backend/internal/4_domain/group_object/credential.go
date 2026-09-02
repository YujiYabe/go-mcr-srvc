package group_object

import (
	"fmt"

	typeObject "backend/internal/4_domain/type_object"
)

type Credential struct {
	clientID     typeObject.ClientID
	clientSecret typeObject.ClientSecret
}

type NewCredentialArgs struct {
	ClientID     *string
	ClientSecret *string
}

func (receiver Credential) ClientID() (
	clientID typeObject.ClientID,
) {
	clientID = receiver.clientID
	return
}

func (receiver Credential) ClientSecret() (
	clientSecret typeObject.ClientSecret,
) {
	clientSecret = receiver.clientSecret
	return
}

func (receiver Credential) CanAuthenticate() (
	canAuthenticate bool,
) {
	canAuthenticate = receiver.ClientID().GetValue() != "" && receiver.ClientSecret().GetValue() != ""
	return
}

func (receiver Credential) EnsureReadyToAuthenticate() (
	err error,
) {
	if receiver.ClientID().GetValue() == "" {
		err = fmt.Errorf("client id is required")
		return
	}
	if receiver.ClientSecret().GetValue() == "" {
		err = fmt.Errorf("client secret is required")
		return
	}

	err = nil
	return
}

func (receiver *Credential) RotateSecret(
	value *string,
) (
	err error,
) {
	clientSecret, err := typeObject.NewClientSecret(value)
	if err != nil {
		return
	}

	receiver.clientSecret = clientSecret

	err = nil
	return
}

func NewCredential(
	args *NewCredentialArgs,
) (
	credential *Credential,
	err error,
) {
	credential = &Credential{}
	if args == nil {
		args = &NewCredentialArgs{}
	}

	credential.clientID, err = typeObject.NewClientID(
		args.ClientID,
	)
	if err != nil {
		credential = nil
		return
	}

	credential.clientSecret, err = typeObject.NewClientSecret(args.ClientSecret)
	if err != nil {
		credential = nil
		return
	}

	return
}
