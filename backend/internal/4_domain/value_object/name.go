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
		name.SetError(name.content.GetError())
	}

	return
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
