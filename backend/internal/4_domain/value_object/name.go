package value_object

import (
	primitiveObject "backend/internal/4_domain/primitive_object"
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
	value *string,
) (
	name Name,
) {
	name = Name{}
	name.SetValue(value)

	return
}

func (receiver *Name) SetValue(
	value *string,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(nameLengthMax),
		primitiveString.WithMinLength(nameLengthMin),
		primitiveString.WithCheckSpell(nameCheckSpell),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		receiver.SetError(receiver.content.GetError())
	}
}

func (receiver *Name) GetError() error {
	return receiver.err
}

func (receiver *Name) SetError(
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
