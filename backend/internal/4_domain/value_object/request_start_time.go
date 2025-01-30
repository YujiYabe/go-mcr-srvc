package value_object

import (
	"context"
	"time"

	primitiveObject "backend/internal/4_domain/primitive_object"
)

const (
	RequestStartTimeMetaName    primitiveObject.ContextKey = "request-start-time"
	RequestStartTimeContextName primitiveObject.ContextKey = "requestStartTime"
)

var (
	requestStartTimeMaxDigit uint = 20
	requestStartTimeMinDigit uint = 0
)

type RequestStartTime struct {
	err     error
	content *primitiveObject.PrimitiveIntX[int64]
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
	primitiveIntX := &primitiveObject.PrimitiveIntX[int64]{}

	if value == nil {
		// デフォルト値を設定
		now := time.Now().UnixMilli()
		value = &now
	}

	receiver.content = primitiveObject.NewPrimitiveIntX(
		primitiveIntX.WithValue(value),
		primitiveIntX.WithMaxDigit(&requestStartTimeMaxDigit),
		primitiveIntX.WithMinDigit(&requestStartTimeMinDigit),
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
}

func (receiver *RequestStartTime) GetValue() int64 {
	return receiver.content.GetValue()
}

func (receiver *RequestStartTime) GetString() string {
	return receiver.content.GetString()
}
