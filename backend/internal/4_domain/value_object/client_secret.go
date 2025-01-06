package value_object

import (
	primitiveObject "backend/internal/4_domain/primitive_object"
)

const (
	clientSecretLengthMax = 99999999999
	clientSecretLengthMin = 0
)

type ClientSecret struct {
	Err     error
	Content *primitiveObject.PrimitiveString
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

	clientSecret.Content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(clientSecretLengthMax),
		primitiveString.WithMinLength(clientSecretLengthMin),
	)

	clientSecret.Content.Validation()
	if clientSecret.Content.GetError() != nil {
		clientSecret.SetError(clientSecret.Content.GetError())
	}

	return
}
func (receiver *ClientSecret) GetError() error {
	return receiver.Err
}

func (receiver *ClientSecret) SetError(
	err error,
) {
	receiver.Err = err
}
