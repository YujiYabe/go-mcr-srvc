package value_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
)

const (
	ClientIPMetaName    primitiveObject.ContextKey = "client-ip"
	ClientIPContextName primitiveObject.ContextKey = "clientIP"
)

var (
	clientIPMaxLength uint = 30
	clientIPMinLength uint = 0
)

type ClientIP struct {
	err     error
	content *primitiveObject.PrimitiveString
}

func NewClientIP(
	ctx context.Context,
	value *string,
) (
	clientIP ClientIP,
) {
	clientIP = ClientIP{}
	clientIP.SetValue(ctx, value)

	return
}

func (receiver *ClientIP) SetValue(
	ctx context.Context,
	value *string,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&clientIPMaxLength),
		primitiveString.WithMinLength(&clientIPMinLength),
	)

	if receiver.content.GetError() != nil {
		receiver.SetError(
			ctx,
			receiver.content.GetError(),
		)
	}
}
func (receiver *ClientIP) GetError() error {
	return receiver.err
}

func (receiver *ClientIP) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
}

func (receiver *ClientIP) GetValue() string {
	return receiver.content.GetValue()
}
