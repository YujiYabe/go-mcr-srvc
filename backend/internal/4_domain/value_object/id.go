package value_object

import (
	primitiveObject "backend/internal/4_domain/primitive_object"
)

const (
	idLengthMax = 99999999999
	idLengthMin = 0
)

type ID struct {
	Err     error
	Content *primitiveObject.PrimitiveInt
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
	id.Content = primitiveObject.NewPrimitiveInt(
		primitiveInt.WithValue(valueInt),
		primitiveInt.WithIsNil(isNil),
		primitiveInt.WithMaxValue(idLengthMax),
		primitiveInt.WithMinValue(idLengthMin),
	)

	id.Content.Validation()
	if id.Content.GetError() != nil {
		id.SetError(id.Content.GetError())
	}

	return
}

func (receiver *ID) GetError() error {
	return receiver.Err
}

func (receiver *ID) SetError(
	err error,
) {
	receiver.Err = err
}
