package primitive_object

import (
	"fmt"
	"sort"
)

type PrimitiveSliceString struct {
	Err       error // バリデーションエラーを格納
	Value     []PrimitiveString
	IsNil     bool // nil状態を示すフラグ
	MaxLength int  // 最大文字列長 (-1は制限なし)
	MinLength int  // 最小文字列長 (-1は制限なし)
}

type PrimitiveSliceStringOption func(*PrimitiveSliceString)

func (receiver *PrimitiveSliceString) WithError(err error) PrimitiveSliceStringOption {
	return func(s *PrimitiveSliceString) {
		s.Err = err
	}
}

func (receiver *PrimitiveSliceString) WithValue(value []PrimitiveString) PrimitiveSliceStringOption {
	return func(s *PrimitiveSliceString) {
		s.Value = value
	}
}

func (receiver *PrimitiveSliceString) WithIsNil(isNil bool) PrimitiveSliceStringOption {
	return func(s *PrimitiveSliceString) {
		s.IsNil = isNil
	}
}

func (receiver *PrimitiveSliceString) WithMaxLength(value int) PrimitiveSliceStringOption {
	return func(s *PrimitiveSliceString) {
		s.MaxLength = value
	}
}

func (receiver *PrimitiveSliceString) WithMinLength(value int) PrimitiveSliceStringOption {
	return func(s *PrimitiveSliceString) {
		s.MinLength = value
	}
}

func NewPrimitiveSliceString(
	options ...PrimitiveSliceStringOption,
) (
	primitiveSliceString *PrimitiveSliceString,
) {
	primitiveSliceString = &PrimitiveSliceString{
		Err:       nil,
		Value:     []PrimitiveString{},
		IsNil:     false,
		MaxLength: -1,
		MinLength: -1,
	}

	for _, option := range options {
		option(primitiveSliceString)
	}

	return
}

func (receiver *PrimitiveSliceString) SetIsNil(isNil bool) {
	receiver.IsNil = isNil
}

func (receiver *PrimitiveSliceString) GetError() error {
	return receiver.Err
}

func (receiver *PrimitiveSliceString) SetError(errString string) {
	receiver.Err = fmt.Errorf("PrimitiveSliceString: %s", errString)
}

func (receiver *PrimitiveSliceString) GetValue() []PrimitiveString {
	if receiver.IsNil {
		receiver.SetError("is nil")
		return []PrimitiveString{}
	}
	return receiver.Value
}

func (receiver *PrimitiveSliceString) SetValue(value []PrimitiveString) {
	if receiver.IsNil {
		receiver.SetError("is nil")
		return
	}
	receiver.Value = value
}

func (receiver *PrimitiveSliceString) SortAsc() {
	sort.Slice(receiver.Value, func(i, j int) bool {
		return receiver.Value[i].Value < receiver.Value[j].Value
	})
}

func (receiver *PrimitiveSliceString) SortDesc() {
	sort.Slice(receiver.Value, func(i, j int) bool {
		return receiver.Value[i].Value > receiver.Value[j].Value
	})
}

func (receiver *PrimitiveSliceString) Validation() error {
	if receiver.IsNil {
		return nil
	}

	if receiver.MaxLength != -1 && len(receiver.Value) > receiver.MaxLength {
		return fmt.Errorf("PrimitiveSliceString: length exceeds maximum allowed (%d)", receiver.MaxLength)
	}

	if receiver.MinLength != -1 && len(receiver.Value) < receiver.MinLength {
		return fmt.Errorf("PrimitiveSliceString: length is less than minimum required (%d)", receiver.MinLength)
	}

	for _, v := range receiver.Value {
		if err := v.Validation(); err != nil {
			return err
		}
	}

	return nil
}
