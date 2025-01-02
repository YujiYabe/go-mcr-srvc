package value_object

import (
	"backend/internal/4_domain/primitive_object"
)

const (
	accessTokenLengthMax = 99999999999
	accessTokenLengthMin = 0
)

type AccessToken struct {
	Content *primitive_object.PrimitiveString
}

func NewAccessToken(
	value *string,
) (
	accessToken AccessToken,
	err error,
) {
	accessToken = AccessToken{}
	primitiveString := &primitive_object.PrimitiveString{}

	valueString, isNil := primitiveString.CheckNil(value)

	accessToken.Content = primitive_object.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(accessTokenLengthMax),
		primitiveString.WithMinLength(accessTokenLengthMin),
	)

	err = accessToken.Content.Validation()

	return
}
