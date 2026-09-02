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
	option = func(s *PrimitiveSliceString) {
		s.value = resValue
		s.isNil = isNil
	}

	return
}

// ______________________________________
func (receiver *PrimitiveSliceString) WithIsNil(
	isNil bool,
) (
	value PrimitiveSliceStringOption,
) {
	value = func(s *PrimitiveSliceString) {
		s.isNil = isNil
	}

	return
}

// ______________________________________
func (receiver *PrimitiveSliceString) WithMaxLength(
	value *uint,
) (
	option PrimitiveSliceStringOption,
) {
	option = func(s *PrimitiveSliceString) {
		s.MaxLength = value
	}

	return
}

// ______________________________________
func (receiver *PrimitiveSliceString) WithMinLength(
	value *uint,
) (
	option PrimitiveSliceStringOption,
) {
	option = func(s *PrimitiveSliceString) {
		s.MinLength = value
	}

	return
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
		primitiveStrings = []PrimitiveString{}

		return
	}
	primitiveStrings = receiver.value

	return
}

// ______________________________________
func (receiver PrimitiveSliceString) GetIsNil() (
	ok bool,
) {
	ok = receiver.isNil

	return
}

// ______________________________________
func (receiver PrimitiveSliceString) Count() (
	value int,
) {
	if receiver.GetIsNil() {
		value = 0

		return
	}

	value = len(receiver.value)

	return
}

// ______________________________________
func (receiver PrimitiveSliceString) IsEmpty() (
	isEmpty bool,
) {
	isEmpty = receiver.Count() == 0

	return
}

// ______________________________________
func (receiver PrimitiveSliceString) HasValue(
	value string,
) (
	hasValue bool,
) {
	if receiver.GetIsNil() {
		hasValue = false

		return
	}

	for _, content := range receiver.value {
		if content.Equal(value) {
			hasValue = true

			return
		}
	}

	hasValue = false

	return
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
		err = nil

		return
	}

	if returnedErr := receiver.ValidationMaxLength(); returnedErr != nil {
		err = returnedErr

		return
	}

	if returnedErr := receiver.ValidationMinLength(); returnedErr != nil {
		err = returnedErr

		return
	}

	for _, value := range receiver.value {
		if returnedErr := value.Validation(); returnedErr != nil {
			err = returnedErr

			return
		}
	}

	err = nil

	return
}

// ______________________________________
func (receiver PrimitiveSliceString) ValidationMaxLength() (
	err error,
) {
	if receiver.MaxLength == nil { // 上限値なし
		err = nil

		return
	}

	if uint(receiver.Count()) > *receiver.MaxLength {
		err = fmt.Errorf(
			"PrimitiveSliceString: length exceeds maximum allowed (%d)",
			*receiver.MaxLength,
		)

		return
	}

	err = nil

	return
}

// ______________________________________
func (receiver PrimitiveSliceString) ValidationMinLength() (
	err error,
) {
	if receiver.MinLength == nil { // 下限値なし
		err = nil

		return
	}

	if uint(receiver.Count()) < *receiver.MinLength {
		err = fmt.Errorf(
			"PrimitiveSliceString: length is less than minimum required (%d)",
			*receiver.MinLength,
		)

		return
	}

	err = nil

	return
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
	values = result

	return
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
	valuePointer = &firstString

	return
}
