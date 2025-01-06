package value_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
	"backend/pkg"
)

const (
	clientSecretLengthMax = 99999999999
	clientSecretLengthMin = 0
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
	// clientSecret.SetValue(ctx, value)
	// return

	primitiveString := &primitiveObject.PrimitiveString{}

	isNil := primitiveString.CheckNil(value)
	valueString := ""
	if !isNil {
		valueString = *value
	}

	clientSecret.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(clientSecretLengthMax),
		primitiveString.WithMinLength(clientSecretLengthMin),
	)

	clientSecret.content.Validation()
	if clientSecret.content.GetError() != nil {
		clientSecret.SetError(ctx,
			clientSecret.content.GetError(),
		)
	}

	return
}

// func (receiver *ClientSecret) SetValue(
// 	ctx context.Context,
// 	value *string,
// ) {
// 	primitiveString := &primitiveObject.PrimitiveString{}

// 	receiver.content = primitiveObject.NewPrimitiveString(
// 		primitiveString.WithValue(value),
// 		primitiveString.WithMaxLength(clientSecretLengthMax),
// 		primitiveString.WithMinLength(clientSecretLengthMin),
// 	)

// 	receiver.content.Validation()
// 	if receiver.content.GetError() != nil {
// 		receiver.SetError(ctx,
// 			receiver.content.GetError(),
// 		)
// 	}
// }

func (receiver *ClientSecret) GetError() error {
	return receiver.err
}

func (receiver *ClientSecret) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
	pkg.Logging(ctx, receiver.GetError())
}

func (receiver *ClientSecret) GetValue() string {
	return receiver.content.GetValue()
}
