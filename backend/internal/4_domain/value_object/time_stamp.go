package value_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
	"backend/pkg"
)

const (
	TimeStampMetaName    primitiveObject.ContextKey = "time-stamp"
	TimeStampContextName primitiveObject.ContextKey = "timeStamp"
)

const (
	timeStampLengthMax = 99999999999
	timeStampLengthMin = 0
)

type TimeStamp struct {
	err     error
	content *primitiveObject.PrimitiveString
}

func NewTimeStamp(
	ctx context.Context,
	value *string,
) (
	timeStamp TimeStamp,
) {
	timeStamp = TimeStamp{}
	timeStamp.SetValue(ctx, value)

	return
}

func (receiver *TimeStamp) SetValue(
	ctx context.Context,
	value *string,
) {
	primitiveString := &primitiveObject.PrimitiveString{}
	receiver.content = primitiveObject.NewPrimitiveString(
		// primitiveString.WithValue(value),
		primitiveString.WithMaxLength(timeStampLengthMax),
		primitiveString.WithMinLength(timeStampLengthMin),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		receiver.SetError(
			ctx,
			receiver.content.GetError(),
		)
	}

}

func (receiver *TimeStamp) GetError() error {
	return receiver.err
}

func (receiver *TimeStamp) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
	pkg.Logging(ctx, receiver.err)
}

func (receiver *TimeStamp) GetValue() string {
	return receiver.content.GetValue()
}
