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
	Content *primitiveObject.PrimitiveString
}

func NewTimeStamp(
	value *string,
) (
	timeStamp TimeStamp,
	err error,
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

	err = timeStamp.Content.Validation()

	return
}
