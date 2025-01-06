package value_object

import (
	primitiveObject "backend/internal/4_domain/primitive_object"
)

const (
	RequestStartTimeMetaName    primitiveObject.ContextKey = "request-start-time"
	RequestStartTimeContextName primitiveObject.ContextKey = "requestStartTime"
)

const (
	requestStartTimeLengthMax = 99999999999
	requestStartTimeLengthMin = 0
)

type RequestStartTime struct {
	err     error
	content *primitiveObject.PrimitiveInt64
}

func NewRequestStartTime(
	value *int64,
) (
	requestStartTime RequestStartTime,
) {
	requestStartTime = RequestStartTime{}
	requestStartTime.SetValue(value)

	return
}

func (receiver *RequestStartTime) SetValue(
	value *int64,
) {
	primitiveInt64 := &primitiveObject.PrimitiveInt64{}

	receiver.content = primitiveObject.NewPrimitiveInt64(
		primitiveInt64.WithValue(value),
		primitiveInt64.WithMaxValue(requestStartTimeLengthMax),
		primitiveInt64.WithMinValue(requestStartTimeLengthMin),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		receiver.SetError(receiver.content.GetError())
	}
}

func (receiver *RequestStartTime) GetError() error {
	return receiver.err
}

func (receiver *RequestStartTime) SetError(
	err error,
) {
	receiver.err = err
}

func (receiver *RequestStartTime) GetValue() int64 {
	return receiver.content.GetValue()
}
