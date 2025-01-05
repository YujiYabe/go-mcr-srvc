package value_object

import (
	primitiveObject "backend/internal/4_domain/primitive_object"
)

const (
	clientSecretLengthMax = 99999999999
	clientSecretLengthMin = 0
)

type ClientSecret struct {
	Content *primitiveObject.PrimitiveString
}

func NewClientSecret(
	value *string,
) (
	clientSecret ClientSecret,
	err error,
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

	err = clientSecret.Content.Validation()

	return
}
