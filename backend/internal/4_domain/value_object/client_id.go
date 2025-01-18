package value_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
)

var (
	clientIDMaxLength uint = 99
	clientIDMinLength uint = 0
)

type ClientID struct {
	err     error
	content *primitiveObject.PrimitiveString
}

func NewClientID(
	ctx context.Context,
	value *string,
) (
	clientID ClientID,
) {
	clientID = ClientID{}
	clientID.SetValue(ctx, value)

	return
}

func (receiver *ClientID) SetValue(
	ctx context.Context,
	value *string,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&clientIDMaxLength),
		primitiveString.WithMinLength(&clientIDMinLength),
	)

	if receiver.content.GetError() != nil {
		receiver.SetError(
			ctx,
			receiver.content.GetError(),
		)
	}
}
func (receiver *ClientID) GetError() error {
	return receiver.err
}

func (receiver *ClientID) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
}

func (receiver *ClientID) GetValue() string {
	return receiver.content.GetValue()
}
