package primitive_object

import (
	"fmt"
	"unicode/utf8"
)

type PrimitiveString struct {
	value  *string
	max    *int
	min    *int
	canNil *bool
}

func NewPrimitiveString() *PrimitiveString {
	return &PrimitiveString{}
}

func (receiver *PrimitiveString) GetValue() string {
	if receiver.value != nil {
		return *receiver.value
	}

	return ""
}

func (receiver *PrimitiveString) GetPointer() *string {
	return receiver.value
}

func (receiver *PrimitiveString) SetValue(v *string) (*PrimitiveString, error) {
	receiver.value = v

	// nullがnilであれば判定対象外にヌル
	canNil := receiver.GetCanNil()
	if canNil != nil && // canNilに値が入っており
		!*canNil && // canNilがfalseだが[*canNil == false]
		receiver.value == nil { // valueがnilの場合
		return receiver, fmt.Errorf("error: %s", "validation error")
	}

	if canNil != nil && // canNilに値が入っており
		*canNil && // canNilがtrueだが[*canNil == true
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

func (receiver *PrimitiveString) GetMax() *int {
	return receiver.max
}

func (receiver *PrimitiveString) GetMin() *int {
	return receiver.min
}

func (receiver *PrimitiveString) GetCanNil() *bool {
	return receiver.canNil
}

func (receiver *PrimitiveString) SetMax(v int) *PrimitiveString {
	receiver.max = &v
	return receiver
}

func (receiver *PrimitiveString) SetMin(v int) *PrimitiveString {
	receiver.min = &v
	return receiver
}

func (receiver *PrimitiveString) SetCanNil(v bool) *PrimitiveString {
	receiver.canNil = &v
	return receiver
}

func (receiver *PrimitiveString) ValidationNil() error {
	if receiver.canNil != nil && //canNilがnilでない
		!*receiver.canNil && //canNilがfalse
		receiver.value == nil {
		return fmt.Errorf("error: %s", " is nil")
	}
	return nil
}

func (receiver *PrimitiveString) ValidationMax() error {
	if err := receiver.ValidationNil(); err != nil {
		return err
	}

	if utf8.RuneCountInString(*receiver.value) > *receiver.max {
		return fmt.Errorf("error: %s", "over max limitation")
	}

	return nil
}

func (receiver *PrimitiveString) ValidationMin() error {
	if err := receiver.ValidationNil(); err != nil {
		return err
	}
	if utf8.RuneCountInString(*receiver.value) < *receiver.min {
		return fmt.Errorf("error: %s", "over min limitation")
	}

	return nil
}

func (receiver *PrimitiveString) IsNil() bool {
	return receiver.value == nil
}
