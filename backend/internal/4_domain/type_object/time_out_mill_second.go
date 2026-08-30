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
	timeOutMillSecondMinDigit uint
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
	err = timeOutMillSecond.setValue(value)

	return
}

func (receiver *TimeOutMillSecond) setValue(
	value *int64,
) (
	err error,
) {
	primitiveIntX := &primitiveObject.PrimitiveIntX[int64]{}

	receiver.content = primitiveObject.NewPrimitiveIntX(
		primitiveIntX.WithValue(value),
		primitiveIntX.WithMaxDigit(&timeOutMillSecondMaxDigit),
		primitiveIntX.WithMinDigit(&timeOutMillSecondMinDigit),
	)
	if err := receiver.content.Validation(); err != nil {
		return err
	}
	return nil
}

func (receiver TimeOutMillSecond) GetValue() (
	value int64,
) {
	return receiver.content.GetValue()
}

func (receiver TimeOutMillSecond) GetString() (
	value string,
) {
	return receiver.content.GetString()
}
