package value_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
	"backend/pkg"
)

const (
	LocaleMetaName    primitiveObject.ContextKey = "locale"
	LocaleContextName primitiveObject.ContextKey = "Locale"
)

const (
	localeLengthMax = 100
	localeLengthMin = 0
)

type Locale struct {
	err     error
	content *primitiveObject.PrimitiveString
}

func NewLocale(
	ctx context.Context,
	value *string,
) (
	locale Locale,
) {
	locale = Locale{}
	locale.SetValue(ctx, value)

	return
}

func (receiver *Locale) SetValue(
	ctx context.Context,
	value *string,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	maxLength := uint(localeLengthMax)
	minLength := uint(localeLengthMin)

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
func (receiver *Locale) GetError() error {
	return receiver.err
}

func (receiver *Locale) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
	pkg.Logging(ctx, receiver.GetError())
}

func (receiver *Locale) GetValue() string {
	return receiver.content.GetValue()
}
