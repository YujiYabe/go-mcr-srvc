package value_object

import (
	"context"

	"github.com/google/uuid"

	primitiveObject "backend/internal/4_domain/primitive_object"
)

const (
	TraceIDMetaName    primitiveObject.ContextKey = "trace-id"
	TraceIDContextName primitiveObject.ContextKey = "traceID"
)

var (
	traceIDMaxLength uint = 36 // length of uuid
	traceIDMinLength uint = 36 // length of uuid
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
	if value == nil {
		// デフォルト値を設定
		newUUID := uuid.New().String()
		value = &newUUID
	}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&traceIDMaxLength),
		primitiveString.WithMinLength(&traceIDMinLength),
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
}

func GetTraceID(
	ctx context.Context,
) (
	value string,
) {
	traceID, ok := ctx.Value(TraceIDContextName).(string)
	if ok {
		value = traceID
	}

	return
}
