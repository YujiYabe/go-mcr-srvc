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
	err     error
	content *primitiveObject.PrimitiveString
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
	traceID.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(valueString),
		primitiveString.WithIsNil(isNil),
		primitiveString.WithMaxLength(traceIDLengthMax),
		primitiveString.WithMinLength(traceIDLengthMin),
	)

	traceID.content.Validation()
	if traceID.content.GetError() != nil {
		traceID.SetError(traceID.content.GetError())
	}

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
func (receiver *TraceID) GetError() error {
	return receiver.err
}

func (receiver *TraceID) SetError(
	err error,
) {
	receiver.err = err
}

func (receiver *TraceID) GetValue() string {
	return receiver.content.GetValue()
}
