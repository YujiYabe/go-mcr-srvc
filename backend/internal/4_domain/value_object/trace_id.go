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
) {
	traceID = TraceID{}

	return
}
func (receiver *TraceID) SetValue(
	value *string,
) {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(traceIDLengthMax),
		primitiveString.WithMinLength(traceIDLengthMin),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		receiver.SetError(receiver.content.GetError())
	}

}

func (receiver *TraceID) GetValue() string {
	return receiver.content.GetValue()
}

func (receiver *TraceID) GetError() error {
	return receiver.err
}

func (receiver *TraceID) SetError(
	err error,
) {
	receiver.err = err
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
