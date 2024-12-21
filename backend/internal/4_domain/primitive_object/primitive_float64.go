package primitive_object

import (
	"fmt"
	"math/big"
)

type PrimitiveFloat64 struct {
	value  *float64
	max    *float64
	min    *float64
	canNil *bool
}

func NewPrimitiveFloat64() *PrimitiveFloat64 {
	return &PrimitiveFloat64{}
}

func (receiver *PrimitiveFloat64) GetValue() float64 {
	if receiver.value != nil {
		return *receiver.value
	}

	return 0

}

func (receiver *PrimitiveFloat64) GetPointer() *float64 {
	return receiver.value
}

func (receiver *PrimitiveFloat64) SetValue(v *float64) (*PrimitiveFloat64, error) {
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

	// max があればmax判定確認
	if receiver.GetMax() != nil {
		if err := receiver.ValidationMax(); err != nil {
			return receiver, fmt.Errorf("error: %s", "validation error")
		}
	}

	// min があればmin判定確認
	if receiver.GetMin() != nil {
		if err := receiver.ValidationMin(); err != nil {
			return receiver, fmt.Errorf("error: %s", "validation error")
		}
	}

	return receiver, nil
}

func (receiver *PrimitiveFloat64) GetMax() *float64 {
	return receiver.max
}

func (receiver *PrimitiveFloat64) GetMin() *float64 {
	return receiver.min
}

func (receiver *PrimitiveFloat64) GetCanNil() *bool {
	return receiver.canNil
}

func (receiver *PrimitiveFloat64) SetMax(v float64) *PrimitiveFloat64 {
	receiver.max = &v
	return receiver
}

func (receiver *PrimitiveFloat64) SetMin(v float64) *PrimitiveFloat64 {
	receiver.min = &v
	return receiver
}

func (receiver *PrimitiveFloat64) SetCanNil(v bool) *PrimitiveFloat64 {
	receiver.canNil = &v
	return receiver
}

func (receiver *PrimitiveFloat64) ValidationNil() error {
	if receiver.value == nil {
		return fmt.Errorf("error: %s", " is nil")
	}
	return nil
}

func (receiver *PrimitiveFloat64) ValidationMax() error {
	if err := receiver.ValidationNil(); err != nil {
		return err
	}
	if big.NewFloat(*receiver.value).Cmp(big.NewFloat(*receiver.max)) > 0 {
		return fmt.Errorf("error: %s", "over max limitation")
	}

	return nil
}

func (receiver *PrimitiveFloat64) ValidationMin() error {
	if err := receiver.ValidationNil(); err != nil {
		return err
	}
	if big.NewFloat(*receiver.value).Cmp(big.NewFloat(*receiver.min)) < 0 {
		return fmt.Errorf("error: %s", "over min limitation")
	}

	return nil
}

func (receiver *PrimitiveFloat64) IsNil() bool {
	return receiver.value == nil
}
