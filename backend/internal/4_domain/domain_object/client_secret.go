package domain_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
)

var (
	clientSecretMaxLength uint = 999
	clientSecretMinLength uint = 0
)

type ClientSecret struct {
	err     error
	content *primitiveObject.PrimitiveString
}

func NewClientSecret(
	ctx context.Context,
	value *string,
) (
	clientSecret ClientSecret,
) {
	clientSecret = ClientSecret{}
	clientSecret.SetValue(ctx, value)

	return
}

func (receiver *ClientSecret) SetValue(
	ctx context.Context,
	value *string,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	minLength := uint(clientSecretMinLength)
	maxLength := uint(clientSecretMaxLength)

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&maxLength),
		primitiveString.WithMinLength(&minLength),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		receiver.SetError(ctx,
			receiver.content.GetError(),
		)
	}
}
func (receiver *ClientSecret) GetError() error {
	return receiver.err
}

func (receiver *ClientSecret) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
}

func (receiver *ClientSecret) GetValue() string {
	return receiver.content.GetValue()
}
