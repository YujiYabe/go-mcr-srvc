package type_object

import (
	"fmt"

	primitiveObject "backend/internal/4_domain/primitive_object"
)

const (
	AccessTokenHeaderName  primitiveObject.ContextKey = "access-token"
	AccessTokenContextName primitiveObject.ContextKey = "AccessToken"
)

var (
	accessTokenMaxLength uint = 9999
	accessTokenMinLength uint
)

type AccessToken struct {
	content *primitiveObject.PrimitiveString
}

func NewAccessToken(
	value *string,
) (
	accessToken AccessToken,
	err error,
) {
	accessToken = AccessToken{}
	err = accessToken.setValue(value)

	return
}

func (receiver *AccessToken) setValue(
	value *string,
) (
	err error,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&accessTokenMaxLength),
		primitiveString.WithMinLength(&accessTokenMinLength),
	)
	if returnedErr := receiver.content.Validation(); returnedErr != nil {
		err = returnedErr

		return
	}
	err = nil

	return
}

func (receiver *AccessToken) ErrorString(
	errString string,
) (
	err error,
) {
	err = fmt.Errorf("error: %s", errString)

	return
}

func (receiver AccessToken) GetValue() (
	value string,
) {
	value = receiver.content.GetValue()

	return
}
