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
	Err     error
	Content *primitiveObject.PrimitiveString
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

	name.Content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(nameLengthMax),
		primitiveString.WithMinLength(nameLengthMin),
		primitiveString.WithCheckSpell(nameCheckSpell),
	)

	name.Content.Validation()
	if name.Content.GetError() != nil {
		name.SetError(name.Content.GetError())
	}

	return
}
func (receiver *Name) GetError() error {
	return receiver.Err
}

func (receiver *Name) SetError(
	err error,
) {
	receiver.Err = err
}
