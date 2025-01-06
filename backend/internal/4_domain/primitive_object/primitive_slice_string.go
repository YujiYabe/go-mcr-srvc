package primitive_object

import (
	"fmt"
	"sort"
)

type PrimitiveSliceString struct {
	err       error // バリデーションエラーを格納
	value     []PrimitiveString
	isNil     bool // nil状態を示すフラグ
	MaxLength int  // 最大列長 (-1は制限なし)
	MinLength int  // 最小列長 (-1は制限なし)
}

type PrimitiveSliceStringOption func(*PrimitiveSliceString)

func (receiver *PrimitiveSliceString) WithError(
	err error,
) PrimitiveSliceStringOption {
	return func(s *PrimitiveSliceString) {
		s.err = err
	}
}

func (receiver *PrimitiveSliceString) WithValue(
	value []PrimitiveString,
) PrimitiveSliceStringOption {
	return func(s *PrimitiveSliceString) {
		s.value = value
	}
}

func (receiver *PrimitiveSliceString) WithIsNil(
	isNil bool,
) PrimitiveSliceStringOption {
	return func(s *PrimitiveSliceString) {
		s.isNil = isNil
	}
}

func (receiver *PrimitiveSliceString) WithMaxLength(
	value int,
) PrimitiveSliceStringOption {
	return func(s *PrimitiveSliceString) {
		s.MaxLength = value
	}
}

func (receiver *PrimitiveSliceString) WithMinLength(
	value int,
) PrimitiveSliceStringOption {
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
		err:       nil,
		value:     []PrimitiveString{},
		isNil:     false,
		MaxLength: -1,
		MinLength: -1,
	}

	for _, option := range options {
		option(primitiveSliceString)
	}

	return
}

func (receiver *PrimitiveSliceString) SetIsNil(
	isNil bool,
) {
	receiver.isNil = isNil
}

func (receiver *PrimitiveSliceString) GetError() error {
	return receiver.err
}

func (receiver *PrimitiveSliceString) SetError(
	err error,
) {
	receiver.err = err
}

func (receiver *PrimitiveSliceString) SetErrorString(
	errString string,
) {
	receiver.SetError(
		fmt.Errorf(
			"error: %s",
			errString,
		),
	)
}

func (receiver *PrimitiveSliceString) GetValue() []PrimitiveString {
	if receiver.isNil {
		receiver.SetErrorString("is nil")
		return []PrimitiveString{}
	}
	return receiver.value
}

func (receiver *PrimitiveSliceString) SetValue(
	valueList []PrimitiveString,
) {
	if valueList == nil {
		receiver.SetIsNil(true)
		return
	}
	receiver.SetIsNil(false)
	receiver.value = valueList
}

func (receiver *PrimitiveSliceString) SortAsc() {
	sort.Slice(receiver.value, func(i, j int) bool {
		return receiver.value[i].value < receiver.value[j].value
	})
}

func (receiver *PrimitiveSliceString) SortDesc() {
	sort.Slice(receiver.value, func(i, j int) bool {
		return receiver.value[i].value > receiver.value[j].value
	})
}

func (receiver *PrimitiveSliceString) Validation() error {
	if receiver.isNil {
		return nil
	}

	if receiver.MaxLength != -1 && len(receiver.value) > receiver.MaxLength {
		return fmt.Errorf(
			"PrimitiveSliceString: length exceeds maximum allowed (%d)",
			receiver.MaxLength,
		)
	}

	if receiver.MinLength != -1 && len(receiver.value) < receiver.MinLength {
		return fmt.Errorf(
			"PrimitiveSliceString: length is less than minimum required (%d)",
			receiver.MinLength,
		)
	}

	for _, value := range receiver.value {
		value.Validation()
		if value.GetError() != nil {
			receiver.SetError(value.GetError())
			break
		}
	}

	return nil
}
