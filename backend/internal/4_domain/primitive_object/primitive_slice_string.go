package primitive_object

import (
	"fmt"
	"sort"
)

// ______________________________________
type PrimitiveSliceString struct {
	err       error // バリデーションエラーを格納
	value     []PrimitiveString
	isNil     bool // nil状態を示すフラグ
	MaxLength int  // 最大列長 (-1は制限なし)
	MinLength int  // 最小列長 (-1は制限なし)
}

// ______________________________________
type PrimitiveSliceStringOption func(*PrimitiveSliceString)

// ______________________________________
func (receiver *PrimitiveSliceString) WithError(
	err error,
) PrimitiveSliceStringOption {
	return func(s *PrimitiveSliceString) {
		s.err = err
	}
}

// ______________________________________
func (receiver *PrimitiveSliceString) WithValue(
	value []PrimitiveString,
) PrimitiveSliceStringOption {
	receiver.SetIsNil(true)
	var resValue []PrimitiveString
	if value != nil {
		receiver.SetIsNil(false)
		resValue = value
	}
	return func(s *PrimitiveSliceString) {
		s.value = resValue
	}
}

// ______________________________________
func (receiver *PrimitiveSliceString) WithIsNil(
	isNil bool,
) PrimitiveSliceStringOption {
	return func(s *PrimitiveSliceString) {
		s.isNil = isNil
	}
}

// ______________________________________
func (receiver *PrimitiveSliceString) WithMaxLength(
	value int,
) PrimitiveSliceStringOption {
	return func(s *PrimitiveSliceString) {
		s.MaxLength = value
	}
}

// ______________________________________
func (receiver *PrimitiveSliceString) WithMinLength(
	value int,
) PrimitiveSliceStringOption {
	return func(s *PrimitiveSliceString) {
		s.MinLength = value
	}
}

// ______________________________________
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

// ______________________________________
func (receiver *PrimitiveSliceString) SetIsNil(
	isNil bool,
) {
	receiver.isNil = isNil
}

// ______________________________________
func (receiver *PrimitiveSliceString) GetError() error {
	return receiver.err
}

// ______________________________________
func (receiver *PrimitiveSliceString) SetError(
	err error,
) {
	receiver.err = err
}

// ______________________________________
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

// ______________________________________
func (receiver *PrimitiveSliceString) GetValue() []PrimitiveString {
	if receiver.isNil {
		receiver.SetErrorString("is nil")
		return []PrimitiveString{}
	}
	return receiver.value
}

// ______________________________________
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

// ______________________________________
func (receiver *PrimitiveSliceString) SortAsc() {
	sort.Slice(receiver.value, func(i, j int) bool {
		return receiver.value[i].value < receiver.value[j].value
	})
}

// ______________________________________
func (receiver *PrimitiveSliceString) SortDesc() {
	sort.Slice(receiver.value, func(i, j int) bool {
		return receiver.value[i].value > receiver.value[j].value
	})
}

// ______________________________________
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

// ______________________________________
func ExtractFirstIndexFromSliceString(
	value []string,
) *string {
	firstString := ""
	if len(value) != 0 {
		firstString = value[0]
	}
	return &firstString
}
