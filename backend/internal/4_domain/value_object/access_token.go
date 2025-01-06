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
	Err     error
	Content *primitiveObject.PrimitiveString
}

func NewAccessToken(
	value *string,
) (
	accessToken AccessToken,
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

	accessToken.Content.Validation()
	if accessToken.Content.GetError() != nil {
		accessToken.SetError(accessToken.Content.GetError())
	}

	return
}

func (receiver *AccessToken) GetError() error {
	return receiver.Err
}

func (receiver *AccessToken) SetError(
	err error,
) {
	receiver.Err = err
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
