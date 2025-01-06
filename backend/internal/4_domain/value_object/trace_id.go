package value_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
)

const (
	TraceIDMetaName    primitiveObject.ContextKey = "trace-id"
	TraceIDContextName primitiveObject.ContextKey = "traceID"
)

const (
	traceIDLengthMax = 99999999999
	traceIDLengthMin = 0
)

type TraceID struct {
	Content *primitiveObject.PrimitiveString
}

func NewTraceID(
	value *string,
) (
	traceID TraceID,
	err error,
) {
	traceID = TraceID{}
	primitiveString := &primitiveObject.PrimitiveString{}

	isNil := primitiveString.CheckNil(value)
	valueString := ""
	if !isNil {
		valueString = *value
	}
	traceID.Content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(traceIDLengthMax),
		primitiveString.WithMinLength(traceIDLengthMin),
	)

	err = traceID.Content.Validation()

	return
}

func GetTraceID(
	ctx context.Context,
) (
	traceIDString string,
) {
	traceID, ok := ctx.Value(TraceIDContextName).(string)
	if ok {
		traceIDString = traceID
	}

	return
}
