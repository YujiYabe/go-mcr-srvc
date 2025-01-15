package value_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
	"backend/pkg"
)

const (
	TimeZoneMetaName    primitiveObject.ContextKey = "time-zone"
	TimeZoneContextName primitiveObject.ContextKey = "TimeZone"
)

const (
	timeZoneLengthMax = 9
	timeZoneLengthMin = 0
)

type TimeZone struct {
	err     error
	content *primitiveObject.PrimitiveString
}

func NewTimeZone(
	ctx context.Context,
	value *string,
) (
	timeZone TimeZone,
) {
	timeZone = TimeZone{}
	timeZone.SetValue(ctx, value)

	return
}

func (receiver *TimeZone) SetValue(
	ctx context.Context,
	value *string,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	minLength := uint(timeZoneLengthMin)
	maxLength := uint(timeZoneLengthMax)

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),

		primitiveString.WithMaxLength(&maxLength),
		primitiveString.WithMinLength(&minLength),
	)

	if receiver.content.GetError() != nil {
		receiver.SetError(
			ctx,
			receiver.content.GetError(),
		)
	}
}
func (receiver *TimeZone) GetError() error {
	return receiver.err
}

func (receiver *TimeZone) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
	pkg.Logging(ctx, receiver.GetError())
}

func (receiver *TimeZone) GetValue() string {
	return receiver.content.GetValue()
}
