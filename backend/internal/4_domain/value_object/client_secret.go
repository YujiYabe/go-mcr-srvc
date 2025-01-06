package value_object

import (
	primitiveObject "backend/internal/4_domain/primitive_object"
)

const (
	clientSecretLengthMax = 99999999999
	clientSecretLengthMin = 0
)

type ClientSecret struct {
	err     error
	content *primitiveObject.PrimitiveString
}

func NewClientSecret(
	value *string,
) (
	clientSecret ClientSecret,
) {
	clientSecret = ClientSecret{}
	clientSecret.SetValue(value)

	return
}

func (receiver *ClientSecret) SetValue(
	value *string,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(clientSecretLengthMax),
		primitiveString.WithMinLength(clientSecretLengthMin),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		receiver.SetError(receiver.content.GetError())
	}
}

func (receiver *ClientSecret) GetError() error {
	return receiver.err
}

func (receiver *ClientSecret) SetError(
	err error,
) {
	receiver.err = err
}

func (receiver *ClientSecret) GetValue() string {
	return receiver.content.GetValue()
}
