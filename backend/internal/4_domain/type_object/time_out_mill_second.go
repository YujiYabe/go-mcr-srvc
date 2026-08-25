package type_object

import primitiveObject "backend/internal/4_domain/primitive_object"

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
	content *primitiveObject.PrimitiveIntX[int64]
}

func NewTimeOutMillSecond(
	value *int64,
) (
	timeOutMillSecond TimeOutMillSecond,
	err error,
) {
	timeOutMillSecond = TimeOutMillSecond{}
	err = timeOutMillSecond.SetValue(value)

	return
}

func (receiver *TimeOutMillSecond) SetValue(
	value *int64,
) error {
	primitiveIntX := &primitiveObject.PrimitiveIntX[int64]{}

	receiver.content = primitiveObject.NewPrimitiveIntX(
		primitiveIntX.WithValue(value),
		primitiveIntX.WithMaxDigit(&timeOutMillSecondMaxDigit),
		primitiveIntX.WithMinDigit(&timeOutMillSecondMinDigit),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		return receiver.content.GetError()
	}
	return nil
}

func (receiver *TimeOutMillSecond) GetValue() int64 {
	return receiver.content.GetValue()
}

func (receiver *TimeOutMillSecond) GetString() string {
	return receiver.content.GetString()
}
