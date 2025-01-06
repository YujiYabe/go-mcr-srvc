package value_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
	"backend/pkg"
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
	ctx context.Context,
	value *int,
) (
	id ID,
) {
	id = ID{}
	// id.SetValue(ctx, value)
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
		id.SetError(
			ctx,
			id.content.GetError(),
		)
	}

	return
}

// func (receiver *ID) SetValue(
// 	ctx context.Context,
// 	value *int,
// ) {
// 	primitiveInt := &primitiveObject.PrimitiveInt{}

// 	receiver.content = primitiveObject.NewPrimitiveInt(
// 		primitiveInt.WithValue(value),
// 		primitiveInt.WithMaxValue(idLengthMax),
// 		primitiveInt.WithMinValue(idLengthMin),
// 	)

// 	receiver.content.Validation()
// 	if receiver.content.GetError() != nil {
// 		receiver.SetError(
// 			ctx,
// 			receiver.content.GetError(),
// 		)
// 	}
// }

func (receiver *ID) GetError() error {
	return receiver.err
}

func (receiver *ID) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
	pkg.Logging(ctx, receiver.GetError())
}

func (receiver *ID) GetValue() int {
	return receiver.content.GetValue()
}
