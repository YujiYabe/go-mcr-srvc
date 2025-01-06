package primitive_object

import (
	"fmt"
	"time"
)

type PrimitiveTime struct {
	err      error
	value    time.Time
	isNil    bool
	MaxValue time.Time
	MinValue time.Time
}

type PrimitiveTimeOption func(*PrimitiveTime)

func (receiver *PrimitiveTime) WithError(
	err error,
) PrimitiveTimeOption {
	return func(s *PrimitiveTime) {
		s.err = err
	}
}

func (receiver *PrimitiveTime) WithValue(
	value time.Time,
) PrimitiveTimeOption {
	return func(s *PrimitiveTime) {
		s.value = value
	}
}

func (receiver *PrimitiveTime) WithIsNil(
	isNil bool,
) PrimitiveTimeOption {
	return func(s *PrimitiveTime) {
		s.isNil = isNil
	}
}

func (receiver *PrimitiveTime) WithMaxValue(
	maxTime time.Time,
) PrimitiveTimeOption {
	return func(s *PrimitiveTime) {
		s.MaxValue = maxTime
	}
}

func (receiver *PrimitiveTime) WithMinValue(
	minTime time.Time,
) PrimitiveTimeOption {
	return func(s *PrimitiveTime) {
		s.MinValue = minTime
	}
}

func NewPrimitiveTime(
	options ...PrimitiveTimeOption,
) (
	primitiveTime *PrimitiveTime,
) {
	primitiveTime = &PrimitiveTime{
		err:      nil,
		value:    time.Time{},
		isNil:    false,
		MaxValue: time.Time{},
		MinValue: time.Time{},
	}

	for _, option := range options {
		option(primitiveTime)
	}

	return
}

func (receiver *PrimitiveTime) SetIsNil(isNil bool) {
	receiver.isNil = isNil
}

// --------------------------------------
func (receiver *PrimitiveTime) GetIsNil() bool {
	return receiver.isNil
}

func (receiver *PrimitiveTime) GetError() error {
	return receiver.err
}

func (receiver *PrimitiveTime) SetError(
	err error,
) {
	receiver.err = err
}

func (receiver *PrimitiveTime) SetErrorString(
	errString string,
) {
	receiver.SetError(
		fmt.Errorf(
			"error: %s",
			errString,
		),
	)
}

func (receiver *PrimitiveTime) GetValue() time.Time {
	if receiver.isNil {
		receiver.SetErrorString("is nil")
		return time.Time{}
	}
	return receiver.value
}

func (receiver *PrimitiveTime) SetValue(
	value *time.Time,
) {
	if value == nil {
		receiver.SetIsNil(true)
		return
	}
	receiver.SetIsNil(false)
	receiver.value = *value
}

func (receiver *PrimitiveTime) ValidateMaxValue() {
	if !receiver.MaxValue.IsZero() && receiver.value.After(receiver.MaxValue) {
		receiver.err = fmt.Errorf("PrimitiveTime: value exceeds maximum allowed time")
	}
}

func (receiver *PrimitiveTime) ValidateMinValue() {
	if !receiver.MinValue.IsZero() && receiver.value.Before(receiver.MinValue) {
		receiver.err = fmt.Errorf("PrimitiveTime: value is before minimum allowed time")
	}
}

func (receiver *PrimitiveTime) Validation() {
	if receiver.isNil {
		return
	}

	receiver.ValidateMaxValue()
	if receiver.err != nil {
		return
	}

	receiver.ValidateMinValue()
	if receiver.err != nil {
		return
	}

}

func (receiver *PrimitiveTime) CheckNil(
	value *time.Time,
) (
	isNil bool,
) {
	isNil = true
	if value != nil {
		isNil = false
	}

	return
}
