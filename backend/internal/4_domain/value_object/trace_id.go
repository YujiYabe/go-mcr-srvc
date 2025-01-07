package value_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
	"backend/pkg"
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
	ctx context.Context,
	value *string,
) (
	traceID TraceID,
) {
	traceID = TraceID{}
	traceID.SetValue(ctx, value)

	return
}
func (receiver *TraceID) SetValue(
	ctx context.Context,
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
		receiver.SetError(ctx, receiver.content.GetError())
	}
}

func (receiver *TraceID) GetValue() string {
	return receiver.content.GetValue()
}

func (receiver *TraceID) GetError() error {
	return receiver.err
}

func (receiver *TraceID) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
	pkg.Logging(ctx, receiver.err)
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
