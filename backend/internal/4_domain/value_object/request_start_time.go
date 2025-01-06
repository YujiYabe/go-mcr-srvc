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
	primitiveInt64 := &primitiveObject.PrimitiveInt64{}

	isNil := primitiveInt64.CheckNil(value)
	var valueInt64 int64 = 0
	if !isNil {
		valueInt64 = *value
	}
	requestStartTime.content = primitiveObject.NewPrimitiveInt64(
		primitiveInt64.WithValue(valueInt64),
		primitiveInt64.WithIsNil(isNil),
		primitiveInt64.WithMaxValue(requestStartTimeLengthMax),
		primitiveInt64.WithMinValue(requestStartTimeLengthMin),
	)

	requestStartTime.content.Validation()
	if requestStartTime.content.GetError() != nil {
		requestStartTime.SetError(requestStartTime.content.GetError())
	}

	return
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
