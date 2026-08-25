package type_object

import (
	"time"

	primitiveObject "backend/internal/4_domain/primitive_object"
)

const (
	RequestStartTimeHeaderName  primitiveObject.ContextKey = "request-start-time"
	RequestStartTimeContextName primitiveObject.ContextKey = "requestStartTime"
)

var (
	requestStartTimeMaxDigit uint = 20
	requestStartTimeMinDigit uint = 0
)

type RequestStartTime struct {
	content *primitiveObject.PrimitiveIntX[int64]
}

func NewRequestStartTime(
	value *int64,
) (
	requestStartTime RequestStartTime,
	err error,
) {

	requestStartTime = RequestStartTime{}
	err = requestStartTime.SetValue(value)

	return
}

func (receiver *RequestStartTime) SetValue(
	value *int64,
) error {
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
		return receiver.content.GetError()
	}
	return nil
}

func (receiver *RequestStartTime) GetValue() int64 {
	return receiver.content.GetValue()
}

func (receiver *RequestStartTime) GetString() string {
	return receiver.content.GetString()
}
