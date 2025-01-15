package value_object

import (
	"context"
	"time"

	primitiveObject "backend/internal/4_domain/primitive_object"
	"backend/pkg"
)

const (
	RequestStartTimeMetaName    primitiveObject.ContextKey = "request-start-time"
	RequestStartTimeContextName primitiveObject.ContextKey = "requestStartTime"
)

const (
	requestStartTimeMaxDigit = 20
	requestStartTimeMinDigit = 0
)

type RequestStartTime struct {
	err     error
	content *primitiveObject.PrimitiveInt64
}

func NewRequestStartTime(
	ctx context.Context,
	value *int64,
) (
	requestStartTime RequestStartTime,
) {

	requestStartTime = RequestStartTime{}
	requestStartTime.SetValue(ctx, value)

	return
}

func (receiver *RequestStartTime) SetValue(
	ctx context.Context,
	value *int64,
) {
	primitiveInt64 := &primitiveObject.PrimitiveInt64{}

	if value == nil {
		// デフォルト値を設定
		now := time.Now().UnixMilli()
		value = &now
	}

	maxDigit := uint(requestStartTimeMaxDigit)
	minDigit := uint(requestStartTimeMinDigit)

	receiver.content = primitiveObject.NewPrimitiveInt64(
		primitiveInt64.WithValue(value),
		primitiveInt64.WithMaxDigit(&maxDigit),
		primitiveInt64.WithMinDigit(&minDigit),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		receiver.SetError(
			ctx,
			receiver.content.GetError(),
		)
	}
}
func (receiver *RequestStartTime) GetError() error {
	return receiver.err
}

func (receiver *RequestStartTime) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
	pkg.Logging(ctx, receiver.err)
}

func (receiver *RequestStartTime) GetValue() int64 {
	return receiver.content.GetValue()
}

func (receiver *RequestStartTime) GetString() string {
	return receiver.content.GetString()
}
