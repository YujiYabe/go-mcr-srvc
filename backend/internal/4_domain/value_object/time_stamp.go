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
	Err     error
	Content *primitiveObject.PrimitiveString
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
	timeStamp.Content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(timeStampLengthMax),
		primitiveString.WithMinLength(timeStampLengthMin),
	)

	timeStamp.Content.Validation()
	if timeStamp.Content.GetError() != nil {
		timeStamp.SetError(timeStamp.Content.GetError())
	}

	return
}
func (receiver *TimeStamp) GetError() error {
	return receiver.Err
}

func (receiver *TimeStamp) SetError(
	err error,
) {
	receiver.Err = err
}
