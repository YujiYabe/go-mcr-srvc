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
	id.SetValue(value)

	return
}

func (receiver *ID) SetValue(
	value *int,
) {
	primitiveInt := &primitiveObject.PrimitiveInt{}

	receiver.content = primitiveObject.NewPrimitiveInt(
		primitiveInt.WithValue(value),
		primitiveInt.WithMaxValue(idLengthMax),
		primitiveInt.WithMinValue(idLengthMin),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		receiver.SetError(receiver.content.GetError())
	}
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
