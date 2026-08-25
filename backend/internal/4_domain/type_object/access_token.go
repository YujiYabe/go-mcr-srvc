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
	accessTokenMinLength uint = 0
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
	err = accessToken.SetValue(value)

	return
}

func (receiver *AccessToken) SetValue(
	value *string,
) error {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&accessTokenMaxLength),
		primitiveString.WithMinLength(&accessTokenMinLength),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		return receiver.content.GetError()
	}
	return nil
}

func (receiver *AccessToken) ErrorString(
	errString string,
) error {
	return fmt.Errorf("error: %s", errString)
}

func (receiver *AccessToken) GetValue() string {
	return receiver.content.GetValue()
}
