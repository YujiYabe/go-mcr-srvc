package value_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
	"backend/pkg"
)

const (
	idMaxDigit uint = 99
	idMinDigit uint = 0
)

type ID struct {
	err     error
	content *primitiveObject.PrimitiveInt
}

func NewID(
	ctx context.Context,
	value *int,
) (
	id ID,
) {
	id = ID{}
	id.SetValue(ctx, value)

	return
}

func (receiver *ID) SetValue(
	ctx context.Context,
	value *int,
) {
	primitiveInt := &primitiveObject.PrimitiveInt{}

	maxLength := uint(idMaxDigit)
	minLength := uint(idMinDigit)

	receiver.content = primitiveObject.NewPrimitiveInt(
		primitiveInt.WithValue(value),
		primitiveInt.WithMaxDigit(&maxLength),
		primitiveInt.WithMinDigit(&minLength),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		receiver.SetError(
			ctx, receiver.content.GetError(),
		)
	}

}
func (receiver *ID) GetError() error {
	return receiver.err
}

func (receiver *ID) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
	pkg.Logging(ctx, receiver.GetError())
}

func (receiver *ID) GetValue() int {
	return receiver.content.GetValue()
}
