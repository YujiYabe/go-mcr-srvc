package value_object

import (
	"backend/internal/4_domain/primitive_object"
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
	Content *primitive_object.PrimitiveString
}

func NewName(
	value *string,
) (
	name Name,
	err error,
) {
	name = Name{}
	primitiveString := &primitive_object.PrimitiveString{}

	valueString, isNil := primitiveString.CheckNil(value)

	name.Content = primitive_object.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(nameLengthMax),
		primitiveString.WithMinLength(nameLengthMin),
		primitiveString.WithCheckSpell(nameCheckSpell),
	)

	err = name.Content.Validation()
	return
}
