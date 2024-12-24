package value_object

import (
	"backend/internal/4_domain/primitive_object"
)

const (
	idLengthMax = 99999999999
	idLengthMin = 0
)

type ID struct {
	Content *primitive_object.PrimitiveInt
}

func NewID(
	value *int,
) (
	id ID,
	err error,
) {
	id = ID{}
	primitiveInt := &primitive_object.PrimitiveInt{}

	valueInt, isNil := primitiveInt.CheckNil(value)

	id.Content = primitive_object.NewPrimitiveInt(
		primitiveInt.WithValue(valueInt),
		primitiveInt.WithIsNil(isNil),
		primitiveInt.WithMaxLength(idLengthMax),
		primitiveInt.WithMinLength(idLengthMin),
	)

	err = id.Content.Validation()

	return
}
