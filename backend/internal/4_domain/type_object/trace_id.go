package type_object

import (
	"github.com/google/uuid"

	primitiveObject "backend/internal/4_domain/primitive_object"
)

const (
	TraceIDHeaderName  primitiveObject.ContextKey = "trace-id"
	TraceIDContextName primitiveObject.ContextKey = "traceID"
)

var (
	traceIDMaxLength uint = 36 // length of uuid
	traceIDMinLength uint = 36 // length of uuid
)

type TraceID struct {
	content *primitiveObject.PrimitiveString
}

func NewTraceID(
	value *string,
) (
	traceID TraceID,
	err error,
) {
	traceID = TraceID{}
	err = traceID.SetValue(value)

	return
}
func (receiver *TraceID) SetValue(
	value *string,
) error {
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
		return receiver.content.GetError()
	}
	return nil
}
func (receiver *TraceID) GetValue() string {
	return receiver.content.GetValue()
}
