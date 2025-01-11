package value_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
	"backend/pkg"
)

const (
	timezoneLengthMax = 99999999999
	timezoneLengthMin = 0
)

type Timezone struct {
	err     error
	content *primitiveObject.PrimitiveString
}

func NewTimezone(
	ctx context.Context,
	value *string,
) (
	timezone Timezone,
) {
	timezone = Timezone{}
	timezone.SetValue(ctx, value)

	return
}

func (receiver *Timezone) SetValue(
	ctx context.Context,
	value *string,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(timezoneLengthMax),
		primitiveString.WithMinLength(timezoneLengthMin),
	)

	if receiver.content.GetError() != nil {
		receiver.SetError(
			ctx,
			receiver.content.GetError(),
		)
	}
}

func (receiver *Timezone) GetError() error {
	return receiver.err
}

func (receiver *Timezone) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
	pkg.Logging(ctx, receiver.GetError())
}

func (receiver *Timezone) GetValue() string {
	return receiver.content.GetValue()
}
