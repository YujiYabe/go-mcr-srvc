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
	primitiveString := &primitiveObject.PrimitiveString{}

	isNil := primitiveString.CheckNil(value)
	valueString := ""
	if !isNil {
		valueString = *value
	}

	clientSecret.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(clientSecretLengthMax),
		primitiveString.WithMinLength(clientSecretLengthMin),
	)

	clientSecret.content.Validation()
	if clientSecret.content.GetError() != nil {
		clientSecret.SetError(clientSecret.content.GetError())
	}

	return
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
