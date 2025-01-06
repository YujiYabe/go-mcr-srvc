package value_object

import (
	primitiveObject "backend/internal/4_domain/primitive_object"
)

const (
	idLengthMax = 99999999999
	idLengthMin = 0
)

type ID struct {
	err     error
	content *primitiveObject.PrimitiveInt
}

func NewID(
	value *int,
) (
	id ID,
) {
	id = ID{}
	primitiveInt := &primitiveObject.PrimitiveInt{}

	isNil := primitiveInt.CheckNil(value)
	valueInt := 0
	if !isNil {
		valueInt = *value
	}
	id.content = primitiveObject.NewPrimitiveInt(
		primitiveInt.WithValue(valueInt),
		primitiveInt.WithIsNil(isNil),
		primitiveInt.WithMaxValue(idLengthMax),
		primitiveInt.WithMinValue(idLengthMin),
	)

	id.content.Validation()
	if id.content.GetError() != nil {
		id.SetError(id.content.GetError())
	}

	return
}

func (receiver *ID) GetError() error {
	return receiver.err
}

func (receiver *ID) SetError(
	err error,
) {
	receiver.err = err
}

func (receiver *ID) GetValue() int {
	return receiver.content.GetValue()
}
