package type_object

import (
	"context"
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
	err     error
	content *primitiveObject.PrimitiveString
}

func NewAccessToken(
	ctx context.Context,
	value *string,
) (
	accessToken AccessToken,
) {
	accessToken = AccessToken{}
	accessToken.SetValue(ctx, value)

	return
}

func (receiver *AccessToken) SetValue(
	ctx context.Context,
	value *string,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&accessTokenMaxLength),
		primitiveString.WithMinLength(&accessTokenMinLength),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		receiver.SetError(ctx, receiver.content.GetError())
	}
}
func (receiver *AccessToken) GetError() error {
	return receiver.err
}

func (receiver *AccessToken) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
}

func (receiver *AccessToken) SetErrorString(
	ctx context.Context,
	errString string,
) {
	receiver.SetError(
		ctx,
		fmt.Errorf(
			"error: %s",
			errString,
		),
	)
}

func (receiver *AccessToken) GetValue() string {
	return receiver.content.GetValue()
}
