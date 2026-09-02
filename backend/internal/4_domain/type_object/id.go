package type_object

import primitiveObject "backend/internal/4_domain/primitive_object"

var (
	idMaxDigit uint = 9 // 9桁 = 9999999999まで可
	idMinDigit uint
)

type ID struct {
	content *primitiveObject.PrimitiveIntX[int]
}

func NewID(
	value *int,
) (
	id ID,
	err error,
) {
	id = ID{}
	err = id.setValue(value)

	return
}

func (receiver *ID) setValue(
	value *int,
) (
	err error,
) {
	primitiveIntX := &primitiveObject.PrimitiveIntX[int]{}

	receiver.content = primitiveObject.NewPrimitiveIntX(
		primitiveIntX.WithValue(value),
		primitiveIntX.WithMaxDigit(&idMaxDigit),
		primitiveIntX.WithMinDigit(&idMinDigit),
	)
	if returnedErr := receiver.content.Validation(); returnedErr != nil {
		err = returnedErr

		return
	}
	err = nil

	return
}

func (receiver ID) GetValue() (
	value int,
) {
	value = receiver.content.GetValue()

	return
}

func (receiver ID) ToUint32() (
	value uint32,
	err error,
) {
	value, err = receiver.content.ToUint32()

	return
}
