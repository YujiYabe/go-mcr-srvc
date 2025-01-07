package value_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
	"backend/pkg"
)

const (
	clientIDLengthMax = 99999999999
	clientIDLengthMin = 0
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
		primitiveString.WithMaxLength(clientIDLengthMax),
		primitiveString.WithMinLength(clientIDLengthMin),
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
	pkg.Logging(ctx, receiver.GetError())
}

func (receiver *ClientID) GetValue() string {
	return receiver.content.GetValue()
}
