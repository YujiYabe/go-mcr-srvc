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
	return receiver.clientID
}

func (receiver Credential) ClientSecret() (
	clientSecret typeObject.ClientSecret,
) {
	return receiver.clientSecret
}

func (receiver Credential) CanAuthenticate() (
	canAuthenticate bool,
) {
	return receiver.ClientID().GetValue() != "" && receiver.ClientSecret().GetValue() != ""
}

func (receiver Credential) EnsureReadyToAuthenticate() (
	err error,
) {
	if receiver.ClientID().GetValue() == "" {
		return fmt.Errorf("client id is required")
	}
	if receiver.ClientSecret().GetValue() == "" {
		return fmt.Errorf("client secret is required")
	}

	return nil
}

func (receiver *Credential) RotateSecret(
	value *string,
) (
	err error,
) {
	clientSecret, err := typeObject.NewClientSecret(value)
	if err != nil {
		return err
	}

	receiver.clientSecret = clientSecret

	return nil
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
		return nil, err
	}

	credential.clientSecret, err = typeObject.NewClientSecret(args.ClientSecret)
	if err != nil {
		return nil, err
	}

	return
}
