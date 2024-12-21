package primitive_object

import (
	"fmt"
	"time"
)

type PrimitiveTime struct {
	value  *time.Time
	canNil *bool
}

func NewPrimitiveTime() *PrimitiveTime {
	return &PrimitiveTime{}
}

func (receiver *PrimitiveTime) GetValue() time.Time {
	if receiver.value != nil {
		return *receiver.value
	}

	return time.Now()
}

func (receiver *PrimitiveTime) GetPointer() *time.Time {
	return receiver.value
}

func (receiver *PrimitiveTime) SetValue(v *time.Time) (*PrimitiveTime, error) {
	receiver.value = v

	// nullがnilであれば判定対象外にヌル
	canNil := receiver.GetCanNil()
	if canNil != nil && // canNilに値が入っており
		!*canNil && // canNilがfalseだが[*canNil == false]
		receiver.value == nil { // valueがnilの場合
		return receiver, fmt.Errorf("error: %s", "validation error")
	}

	if canNil != nil && // canNilに値が入っており
		*canNil && // canNilがtrueだが[*canNil == true]
		receiver.value == nil { // valueがnilの場合
		return receiver, nil
	}

	if canNil == nil && // canNilに値が入っておらず
		receiver.value == nil { // valueがnilの場合
		return receiver, nil // あえて設定しない要素の為このまま返す
	}

	return receiver, nil
}

func (receiver *PrimitiveTime) GetCanNil() *bool {
	return receiver.canNil
}

func (receiver *PrimitiveTime) SetCanNil(v bool) *PrimitiveTime {
	receiver.canNil = &v
	return receiver
}

func (receiver *PrimitiveTime) ValidationNil() error {
	if receiver.value == nil {
		return fmt.Errorf("error: %s", " is nil")
	}
	return nil
}

func (receiver *PrimitiveTime) IsNil() bool {
	return receiver.value == nil
}

func (receiver *PrimitiveTime) GetNow() time.Time {
	return time.Now()
}
