package value_object

import (
	"backend/internal/4_domain/primitive_object"
)

const (
	traceIDLengthMax = 99999999999
	traceIDLengthMin = 0
)

type TraceID struct {
	Content *primitive_object.PrimitiveString
}

func NewTraceID(
	value *string,
) (
	traceID TraceID,
	err error,
) {
	traceID = TraceID{}
	primitiveString := &primitive_object.PrimitiveString{}

	isNil := primitiveString.CheckNil(value)
	valueString := ""
	if !isNil {
		valueString = *value
	}
	traceID.Content = primitive_object.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(traceIDLengthMax),
		primitiveString.WithMinLength(traceIDLengthMin),
	)

	err = traceID.Content.Validation()

	return
}
