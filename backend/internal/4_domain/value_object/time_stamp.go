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
	primitiveString := &primitiveObject.PrimitiveString{}

	isNil := primitiveString.CheckNil(value)
	valueString := ""
	if !isNil {
		valueString = *value
	}
	timeStamp.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(timeStampLengthMax),
		primitiveString.WithMinLength(timeStampLengthMin),
	)

	timeStamp.content.Validation()
	if timeStamp.content.GetError() != nil {
		timeStamp.SetError(timeStamp.content.GetError())
	}

	return
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
