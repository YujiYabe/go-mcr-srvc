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
	requestStartTimeMinDigit uint
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
	err = requestStartTime.setValue(value)

	return
}

func (receiver *RequestStartTime) setValue(
	value *int64,
) (
	err error,
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
	if returnedErr := receiver.content.Validation(); returnedErr != nil {
		err = returnedErr

		return
	}
	err = nil

	return
}

func (receiver RequestStartTime) GetValue() (
	value int64,
) {
	value = receiver.content.GetValue()

	return
}

func (receiver RequestStartTime) GetString() (
	value string,
) {
	value = receiver.content.GetString()

	return
}
