package value_object

import (
	primitiveObject "backend/internal/4_domain/primitive_object"
)

const (
	accessTokenLengthMax = 99999999999
	accessTokenLengthMin = 0
)

type AccessToken struct {
	Content *primitiveObject.PrimitiveString
}

func NewAccessToken(
	value *string,
) (
	accessToken AccessToken,
	err error,
) {
	accessToken = AccessToken{}
	primitiveString := &primitiveObject.PrimitiveString{}

	isNil := primitiveString.CheckNil(value)
	valueString := ""
	if !isNil {
		valueString = *value
	}
	accessToken.Content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(accessTokenLengthMax),
		primitiveString.WithMinLength(accessTokenLengthMin),
	)

	err = accessToken.Content.Validation()

	return
}
