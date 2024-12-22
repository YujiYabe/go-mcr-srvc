package value_object

import (
	"backend/internal/4_domain/primitive_object"
)

const (
	nameLengthMax = 30
	nameLengthMin = 1
)

type Name struct {
	primitiveString *primitive_object.PrimitiveString
}

func NewName(
	value string,
) (
	name *Name,
	err error,
) {
	name = &Name{
		primitiveString: primitive_object.NewPrimitiveString(
			primitive_object.WithValue(value),
			primitive_object.WithMaxLength(nameLengthMax),
			primitive_object.WithMinLength(nameLengthMin),
		),
	}

	name.primitiveString.Validation()
	if name.primitiveString.GetError() != nil {
		return name, name.primitiveString.GetError()
	}

	return
}
