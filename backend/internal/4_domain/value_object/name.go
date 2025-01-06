package value_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
	"backend/pkg"
)

const (
	nameLengthMax = 30
	nameLengthMin = 1
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
	// name.SetValue(ctx, value)

	primitiveString := &primitiveObject.PrimitiveString{}

	isNil := primitiveString.CheckNil(value)
	valueString := ""
	if !isNil {
		valueString = *value
	}

	name.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(nameLengthMax),
		primitiveString.WithMinLength(nameLengthMin),
		primitiveString.WithCheckSpell(nameCheckSpell),
	)

	name.content.Validation()
	if name.content.GetError() != nil {
		name.SetError(ctx, name.content.GetError())
	}

	return
}

// func (receiver *Name) SetValue(
// 	ctx context.Context,
// 	value *string,
// ) {
// 	primitiveString := &primitiveObject.PrimitiveString{}

// 	receiver.content = primitiveObject.NewPrimitiveString(
// 		primitiveString.WithValue(value),
// 		primitiveString.WithMaxLength(nameLengthMax),
// 		primitiveString.WithMinLength(nameLengthMin),
// 		primitiveString.WithCheckSpell(nameCheckSpell),
// 	)

// 	receiver.content.Validation()
// 	if receiver.content.GetError() != nil {
// 		receiver.SetError(ctx, receiver.content.GetError())
// 	}
// }

func (receiver *Name) GetError() error {
	return receiver.err
}

func (receiver *Name) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
	pkg.Logging(ctx, receiver.GetError())
}

func (receiver *Name) GetValue() string {
	return receiver.content.GetValue()
}

func (receiver *Name) GetIsNil() bool {
	return receiver.content.GetIsNil()
}
