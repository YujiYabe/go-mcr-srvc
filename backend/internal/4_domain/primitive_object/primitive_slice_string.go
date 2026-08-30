package primitive_object

import (
	"fmt"
	"sort"
)

// ______________________________________
type PrimitiveSliceString struct {
	value     []PrimitiveString
	isNil     bool  // nil状態を示すフラグ
	MaxLength *uint // 最大列長
	MinLength *uint // 最小列長
}

// ______________________________________
type PrimitiveSliceStringOption func(*PrimitiveSliceString)

// ______________________________________
func (receiver *PrimitiveSliceString) WithValue(
	value []PrimitiveString,
) (
	option PrimitiveSliceStringOption,
) {
	isNil := true
	var resValue []PrimitiveString
	if value != nil {
		isNil = false
		resValue = value
	}
	return func(s *PrimitiveSliceString) {
		s.value = resValue
		s.isNil = isNil
	}
}

// ______________________________________
func (receiver *PrimitiveSliceString) WithIsNil(
	isNil bool,
) (
	value PrimitiveSliceStringOption,
) {
	return func(s *PrimitiveSliceString) {
		s.isNil = isNil
	}
}

// ______________________________________
func (receiver *PrimitiveSliceString) WithMaxLength(
	value *uint,
) (
	option PrimitiveSliceStringOption,
) {
	return func(s *PrimitiveSliceString) {
		s.MaxLength = value
	}
}

// ______________________________________
func (receiver *PrimitiveSliceString) WithMinLength(
	value *uint,
) (
	option PrimitiveSliceStringOption,
) {
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
		value:     []PrimitiveString{},
		isNil:     true,
		MaxLength: nil,
		MinLength: nil,
	}

	for _, option := range options {
		option(primitiveSliceString)
	}

	return
}

// ______________________________________
func (receiver PrimitiveSliceString) GetValue() (
	primitiveStrings []PrimitiveString,
) {
	if receiver.isNil {
		return []PrimitiveString{}
	}
	return receiver.value
}

// ______________________________________
func (receiver PrimitiveSliceString) GetIsNil() (
	ok bool,
) {
	return receiver.isNil
}

// ______________________________________
func (receiver PrimitiveSliceString) Count() (
	value int,
) {
	if receiver.GetIsNil() {
		return 0
	}

	return len(receiver.value)
}

// ______________________________________
func (receiver PrimitiveSliceString) IsEmpty() (
	isEmpty bool,
) {
	return receiver.Count() == 0
}

// ______________________________________
func (receiver PrimitiveSliceString) HasValue(
	value string,
) (
	hasValue bool,
) {
	if receiver.GetIsNil() {
		return false
	}

	for _, content := range receiver.value {
		if content.Equal(value) {
			return true
		}
	}

	return false
}

// ______________________________________
func (receiver *PrimitiveSliceString) SortAsc() {
	sort.Slice(receiver.value, func(leftIndex, rightIndex int) bool {
		return receiver.value[leftIndex].value < receiver.value[rightIndex].value
	})
}

// ______________________________________
func (receiver *PrimitiveSliceString) SortDesc() {
	sort.Slice(receiver.value, func(leftIndex, rightIndex int) bool {
		return receiver.value[leftIndex].value > receiver.value[rightIndex].value
	})
}

// ______________________________________
func (receiver PrimitiveSliceString) Validation() (
	err error,
) {
	if receiver.isNil {
		return nil
	}

	if err := receiver.ValidationMaxLength(); err != nil {
		return err
	}

	if err := receiver.ValidationMinLength(); err != nil {
		return err
	}

	for _, value := range receiver.value {
		if err := value.Validation(); err != nil {
			return err
		}
	}

	return nil
}

// ______________________________________
func (receiver PrimitiveSliceString) ValidationMaxLength() (
	err error,
) {
	if receiver.MaxLength == nil { // 上限値なし
		return nil
	}

	if uint(receiver.Count()) > *receiver.MaxLength {
		return fmt.Errorf(
			"PrimitiveSliceString: length exceeds maximum allowed (%d)",
			*receiver.MaxLength,
		)
	}

	return nil
}

// ______________________________________
func (receiver PrimitiveSliceString) ValidationMinLength() (
	err error,
) {
	if receiver.MinLength == nil { // 下限値なし
		return nil
	}

	if uint(receiver.Count()) < *receiver.MinLength {
		return fmt.Errorf(
			"PrimitiveSliceString: length is less than minimum required (%d)",
			*receiver.MinLength,
		)
	}

	return nil
}

// ______________________________________
// []PrimitiveString を []string に変換して出力する関数
func (receiver *PrimitiveSliceString) ToSliceString() (
	values []string,
) {
	result := make([]string, len(receiver.value))
	for index, value := range receiver.value {
		result[index] = value.GetValue()
	}
	return result
}

// ______________________________________
func ExtractFirstIndexFromSliceString(
	value []string,
) (
	valuePointer *string,
) {
	firstString := ""
	if len(value) != 0 {
		firstString = value[0]
	}
	return &firstString
}
