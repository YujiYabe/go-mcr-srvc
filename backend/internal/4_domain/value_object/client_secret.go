package value_object

import (
	"backend/internal/4_domain/primitive_object"
)

const (
	clientSecretLengthMax = 99999999999
	clientSecretLengthMin = 0
)

type ClientSecret struct {
	Content *primitive_object.PrimitiveString
}

func NewClientSecret(
	value *string,
) (
	clientSecret ClientSecret,
	err error,
) {
	clientSecret = ClientSecret{}
	primitiveString := &primitive_object.PrimitiveString{}

	valueString, isNil := primitiveString.CheckNil(value)

	clientSecret.Content = primitive_object.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(clientSecretLengthMax),
		primitiveString.WithMinLength(clientSecretLengthMin),
	)

	err = clientSecret.Content.Validation()

	return
}
