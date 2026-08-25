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
	err = id.setValue(value)

	return
}

func (receiver *ID) setValue(
	value *int,
) error {
	primitiveIntX := &primitiveObject.PrimitiveIntX[int]{}

	receiver.content = primitiveObject.NewPrimitiveIntX(
		primitiveIntX.WithValue(value),
		primitiveIntX.WithMaxDigit(&idMaxDigit),
		primitiveIntX.WithMinDigit(&idMinDigit),
	)
	if err := receiver.content.Validation(); err != nil {
		return err
	}
	return nil
}

func (receiver ID) GetValue() int {
	return receiver.content.GetValue()
}
