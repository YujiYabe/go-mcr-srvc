package primitive_object

import (
	"fmt"
	"time"
)

type PrimitiveTime struct {
	Err      error
	Value    time.Time
	IsNil    bool
	MaxValue time.Time
	MinValue time.Time
}

type PrimitiveTimeOption func(*PrimitiveTime)

func (receiver *PrimitiveTime) WithError(
	err error,
) PrimitiveTimeOption {
	return func(s *PrimitiveTime) {
		s.Err = err
	}
}

func (receiver *PrimitiveTime) WithValue(
	value time.Time,
) PrimitiveTimeOption {
	return func(s *PrimitiveTime) {
		s.Value = value
	}
}

func (receiver *PrimitiveTime) WithIsNil(
	isNil bool,
) PrimitiveTimeOption {
	return func(s *PrimitiveTime) {
		s.IsNil = isNil
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
		Err:      nil,
		Value:    time.Time{},
		IsNil:    false,
		MaxValue: time.Time{},
		MinValue: time.Time{},
	}

	for _, option := range options {
		option(primitiveTime)
	}

	return
}

func (receiver *PrimitiveTime) SetIsNil(isNil bool) {
	receiver.IsNil = isNil
}

func (receiver *PrimitiveTime) GetError() error {
	return receiver.Err
}

func (receiver *PrimitiveTime) SetError(
	err error,
) {
	receiver.Err = err
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
	if receiver.IsNil {
		receiver.SetErrorString("is nil")
		return time.Time{}
	}
	return receiver.Value
}

func (receiver *PrimitiveTime) SetValue(value time.Time) {
	if receiver.IsNil {
		receiver.SetErrorString("is nil")
		return
	}
	receiver.Value = value
}

func (receiver *PrimitiveTime) ValidateMaxValue() {
	if !receiver.MaxValue.IsZero() && receiver.Value.After(receiver.MaxValue) {
		receiver.Err = fmt.Errorf("PrimitiveTime: value exceeds maximum allowed time")
	}
}

func (receiver *PrimitiveTime) ValidateMinValue() {
	if !receiver.MinValue.IsZero() && receiver.Value.Before(receiver.MinValue) {
		receiver.Err = fmt.Errorf("PrimitiveTime: value is before minimum allowed time")
	}
}

func (receiver *PrimitiveTime) Validation() {
	if receiver.IsNil {
		return
	}

	receiver.ValidateMaxValue()
	if receiver.Err != nil {
		return
	}

	receiver.ValidateMinValue()
	if receiver.Err != nil {
		return
	}

	return
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
