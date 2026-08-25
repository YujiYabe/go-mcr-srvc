package type_object

import primitiveObject "backend/internal/4_domain/primitive_object"

var (
	idMaxDigit uint = 9 // 9桁 = 9999999999まで可
	idMinDigit uint = 0
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
	err = id.SetValue(value)

	return
}

func (receiver *ID) SetValue(
	value *int,
) error {
	primitiveIntX := &primitiveObject.PrimitiveIntX[int]{}

	receiver.content = primitiveObject.NewPrimitiveIntX(
		primitiveIntX.WithValue(value),
		primitiveIntX.WithMaxDigit(&idMaxDigit),
		primitiveIntX.WithMinDigit(&idMinDigit),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		return receiver.content.GetError()
	}
	return nil
}

func (receiver *ID) GetValue() int {
	return receiver.content.GetValue()
}
