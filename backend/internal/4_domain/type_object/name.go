package type_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
)

var (
	nameMaxLength uint = 30
	nameMinLength uint = 1
)

var nameCheckSpell = []string{
	"盗む",
	"暴力",
}

type Name struct {
	err     error
	content *primitiveObject.PrimitiveString
}

func NewName(
	ctx context.Context,
	value *string,
) (
	name Name,
) {
	name = Name{}
	name.SetValue(ctx, value)

	return
}

func (receiver *Name) SetValue(
	ctx context.Context,
	value *string,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&nameMaxLength),
		primitiveString.WithMinLength(&nameMinLength),
		primitiveString.WithCheckSpell(nameCheckSpell),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		receiver.SetError(ctx, receiver.content.GetError())
	}
}
func (receiver *Name) GetError() error {
	return receiver.err
}

func (receiver *Name) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
}

func (receiver *Name) GetValue() string {
	return receiver.content.GetValue()
}

func (receiver *Name) GetIsNil() bool {
	return receiver.content.GetIsNil()
}
