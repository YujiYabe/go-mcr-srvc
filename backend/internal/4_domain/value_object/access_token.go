package value_object

import (
	primitiveObject "backend/internal/4_domain/primitive_object"
	"fmt"
)

const (
	AccessTokenMetaName    primitiveObject.ContextKey = "access-token"
	AccessTokenContextName primitiveObject.ContextKey = "AccessToken"
)

const (
	accessTokenLengthMax = 99999999999
	accessTokenLengthMin = 0
)

type AccessToken struct {
	err     error
	content *primitiveObject.PrimitiveString
}

func NewAccessToken(
	value *string,
) (
	accessToken AccessToken,
) {
	accessToken = AccessToken{}
	accessToken.SetValue(value)

	return
}

func (receiver *AccessToken) SetValue(
	value *string,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(accessTokenLengthMax),
		primitiveString.WithMinLength(accessTokenLengthMin),
	)

	receiver.content.Validation()
	if receiver.GetError() != nil {
		receiver.SetError(receiver.GetError())
		return
	}
}

func (receiver *AccessToken) GetError() error {
	return receiver.err
}

func (receiver *AccessToken) SetError(
	err error,
) {
	receiver.err = err
}

func (receiver *AccessToken) SetErrorString(
	errString string,
) {
	receiver.SetError(
		fmt.Errorf(
			"error: %s",
			errString,
		),
	)
}

func (receiver *AccessToken) GetValue() string {
	return receiver.content.GetValue()
}
