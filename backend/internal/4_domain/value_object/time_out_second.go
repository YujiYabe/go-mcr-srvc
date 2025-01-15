package value_object

import (
	"context"

	primitiveObject "backend/internal/4_domain/primitive_object"
	"backend/pkg"
)

const (
	TimeOutSecondValue = 10
)

const (
	TimeOutSecondMetaName    primitiveObject.ContextKey = "time-out-second"
	TimeOutSecondContextName primitiveObject.ContextKey = "timeOutSecond"
)

const (
	timeOutSecondMaxDigit = 10
	timeOutSecondMinDigit = 0
)

type TimeOutSecond struct {
	err     error
	content *primitiveObject.PrimitiveInt64
}

func NewTimeOutSecond(
	ctx context.Context,
	value *int64,
) (
	timeOutSecond TimeOutSecond,
) {
	timeOutSecond = TimeOutSecond{}
	timeOutSecond.SetValue(ctx, value)

	return
}

func (receiver *TimeOutSecond) SetValue(
	ctx context.Context,
	value *int64,
) {
	primitiveInt64 := &primitiveObject.PrimitiveInt64{}

	maxDigit := uint(timeOutSecondMaxDigit)
	minDigit := uint(timeOutSecondMinDigit)

	receiver.content = primitiveObject.NewPrimitiveInt64(
		primitiveInt64.WithValue(value),

		primitiveInt64.WithMaxDigit(&maxDigit),
		primitiveInt64.WithMinDigit(&minDigit),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		receiver.SetError(ctx, receiver.content.GetError())
	}
}

func (receiver *TimeOutSecond) GetError() error {
	return receiver.err
}

func (receiver *TimeOutSecond) SetError(
	ctx context.Context,
	err error,
) {
	receiver.err = err
	pkg.Logging(ctx, receiver.err)
}

func (receiver *TimeOutSecond) GetValue() int64 {
	return receiver.content.GetValue()
}

func (receiver *TimeOutSecond) GetString() string {
	return receiver.content.GetString()
}

func GetTimeoutSecond(
	ctx context.Context,
) (
	value int64,
) {
	timeOutSecond, ok := ctx.Value(TimeOutSecondContextName).(int64)

	if ok {
		value = timeOutSecond
	}

	return
}
