package value_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
	"backend/pkg"
)

const (
	userAgentLengthMax = 99999999999
	userAgentLengthMin = 0
)

type UserAgent struct {
	err     error
	content *primitiveObject.PrimitiveString
}

func NewUserAgent(
	ctx context.Context,
	value *string,
) (
	userAgent UserAgent,
) {
	userAgent = UserAgent{}
	userAgent.SetValue(ctx, value)

	return
}

func (receiver *UserAgent) SetValue(
	ctx context.Context,
	value *string,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(userAgentLengthMax),
		primitiveString.WithMinLength(userAgentLengthMin),
	)

	if receiver.content.GetError() != nil {
		receiver.SetError(
			ctx,
			receiver.content.GetError(),
		)
	}
}

func (receiver *UserAgent) GetError() error {
	return receiver.err
}

func (receiver *UserAgent) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
	pkg.Logging(ctx, receiver.GetError())
}

func (receiver *UserAgent) GetValue() string {
	return receiver.content.GetValue()
}
