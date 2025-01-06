package value_object

import (
	primitiveObject "backend/internal/4_domain/primitive_object"
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
	value *string,
) (
	timeStamp TimeStamp,
) {
	timeStamp = TimeStamp{}
	timeStamp.SetValue(value)

	return
}

func (receiver *TimeStamp) SetValue(
	value *string,
) {
	primitiveString := &primitiveObject.PrimitiveString{}
	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(timeStampLengthMax),
		primitiveString.WithMinLength(timeStampLengthMin),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		receiver.SetError(receiver.content.GetError())
	}

}

func (receiver *TimeStamp) GetError() error {
	return receiver.err
}

func (receiver *TimeStamp) SetError(
	err error,
) {
	receiver.err = err
}

func (receiver *TimeStamp) GetValue() string {
	return receiver.content.GetValue()
}
