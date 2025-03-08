package domain_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
)

const (
	TimeOutMillSecondValue = 10
)

const (
	TimeOutMillSecondHeaderName  primitiveObject.ContextKey = "time-out-second"
	TimeOutMillSecondContextName primitiveObject.ContextKey = "timeOutMillSecond"
)

var (
	timeOutMillSecondMaxDigit uint = 10
	timeOutMillSecondMinDigit uint = 0
)

type TimeOutMillSecond struct {
	err     error
	content *primitiveObject.PrimitiveIntX[int64]
}

func NewTimeOutMillSecond(
	ctx context.Context,
	value *int64,
) (
	timeOutMillSecond TimeOutMillSecond,
) {
	timeOutMillSecond = TimeOutMillSecond{}
	timeOutMillSecond.SetValue(ctx, value)

	return
}

func (receiver *TimeOutMillSecond) SetValue(
	ctx context.Context,
	value *int64,
) {
	primitiveIntX := &primitiveObject.PrimitiveIntX[int64]{}

	receiver.content = primitiveObject.NewPrimitiveIntX(
		primitiveIntX.WithValue(value),
		primitiveIntX.WithMaxDigit(&timeOutMillSecondMaxDigit),
		primitiveIntX.WithMinDigit(&timeOutMillSecondMinDigit),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		receiver.SetError(ctx, receiver.content.GetError())
	}
}

func (receiver *TimeOutMillSecond) GetError() error {
	return receiver.err
}

func (receiver *TimeOutMillSecond) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
}

func (receiver *TimeOutMillSecond) GetValue() int64 {
	return receiver.content.GetValue()
}

func (receiver *TimeOutMillSecond) GetString() string {
	return receiver.content.GetString()
}

func GetTimeoutSecond(
	ctx context.Context,
) (
	value int64,
) {
	timeOutMillSecond, ok := ctx.Value(TimeOutMillSecondContextName).(int64)

	if ok {
		value = timeOutMillSecond
	}

	return
}
