package primitive_object

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type ContextKey string

// ______________________________________
// PrimitiveString は文字列値に対してバリデーション機能を提供する構造体です。
// nil チェック、長さ制限、禁止文字列のチェックなどの機能を備えています。
type PrimitiveString struct {
	value     string   // 実際の文字列値
	isNil     bool     // nil状態を示すフラグ
	maxLength *uint    // 最大文字列長
	minLength *uint    // 最小文字列長
	spellList []string // チェック対象の禁止文字列リスト
}

// ______________________________________
// NewPrimitiveString は指定されたオプションで新しいPrimitiveStringインスタンスを生成します
type PrimitiveStringOption func(*PrimitiveString)

// ______________________________________
// WithValue は文字列値を設定するオプションを返します
func (receiver *PrimitiveString) WithValue(
	value *string,
) (
	option PrimitiveStringOption,
) {
	isNil := receiver.CheckNil(value)
	valueString := ""
	if !isNil {
		valueString = *value
	}

	option = func(s *PrimitiveString) {
		s.value = valueString
		s.isNil = isNil
	}

	return
}

// ______________________________________
// WithIsNil はnil状態を設定するオプションを返します
func (receiver *PrimitiveString) WithIsNil(
	isNil bool,
) (
	value PrimitiveStringOption,
) {
	value = func(s *PrimitiveString) {
		s.isNil = isNil
	}

	return
}

// ______________________________________
// WithMaxLength は最大文字列長を設定するオプションを返します
func (receiver *PrimitiveString) WithMaxLength(
	length *uint,
) (
	value PrimitiveStringOption,
) {
	value = func(s *PrimitiveString) {
		s.maxLength = length
	}

	return
}

// ______________________________________
// WithMinLength は最小文字列長を設定するオプションを返します
func (receiver *PrimitiveString) WithMinLength(
	length *uint,
) (
	value PrimitiveStringOption,
) {
	value = func(s *PrimitiveString) {
		s.minLength = length
	}

	return
}

// ______________________________________
// WithCheckSpell は禁止文字列リストを設定するオプションを返します
func (receiver *PrimitiveString) WithCheckSpell(
	spellList []string,
) (
	value PrimitiveStringOption,
) {
	value = func(s *PrimitiveString) {
		s.spellList = spellList
	}

	return
}

// ______________________________________
func NewPrimitiveString(
	options ...PrimitiveStringOption,
) (
	primitiveString *PrimitiveString,
) {

	// デフォルト値を設定
	primitiveString = &PrimitiveString{
		value:     "",
		isNil:     true,
		maxLength: nil,
		minLength: nil,
		spellList: []string{},
	}

	// オプションを適用
	for _, option := range options {
		option(primitiveString)
	}

	return
}

// ______________________________________
func (receiver PrimitiveString) GetIsNil() (
	ok bool,
) {
	ok = receiver.isNil

	return
}

// ______________________________________
func (receiver PrimitiveString) newErrorString(
	errString string,
) (
	err error,
) {
	err = fmt.Errorf(
		"error: %s",
		errString,
	)

	return
}

// ______________________________________
func (receiver PrimitiveString) GetValue() (
	value string,
) {
	if receiver.GetIsNil() {
		value = ""

		return
	}
	value = receiver.value

	return
}

func (receiver PrimitiveString) ToInt() (
	value int,
	err error,
) {
	value = 0
	err = nil
	intValue, parseErr := strconv.ParseInt(receiver.GetValue(), 10, strconv.IntSize)
	if parseErr != nil {
		err = parseErr

		return
	}
	value = int(intValue)

	return
}

func (receiver PrimitiveString) ToInt32() (
	value int32,
	err error,
) {
	value = 0
	err = nil
	intValue, parseErr := strconv.ParseInt(receiver.GetValue(), 10, 32)
	if parseErr != nil {
		err = parseErr

		return
	}
	value = int32(intValue)

	return
}

func (receiver PrimitiveString) ToInt64() (
	value int64,
	err error,
) {
	value, err = strconv.ParseInt(receiver.GetValue(), 10, 64)

	return
}

func (receiver PrimitiveString) ToUint() (
	value uint,
	err error,
) {
	value = 0
	err = nil
	uintValue, parseErr := strconv.ParseUint(receiver.GetValue(), 10, strconv.IntSize)
	if parseErr != nil {
		err = parseErr

		return
	}
	value = uint(uintValue)

	return
}

func (receiver PrimitiveString) ToUint32() (
	value uint32,
	err error,
) {
	value = 0
	err = nil
	uintValue, parseErr := strconv.ParseUint(receiver.GetValue(), 10, 32)
	if parseErr != nil {
		err = parseErr

		return
	}
	value = uint32(uintValue)

	return
}

func (receiver PrimitiveString) ToUint64() (
	value uint64,
	err error,
) {
	value, err = strconv.ParseUint(receiver.GetValue(), 10, 64)

	return
}

// ______________________________________
func (receiver PrimitiveString) Length() (
	value uint,
) {
	if receiver.GetIsNil() {
		value = 0

		return
	}

	value = uint(utf8.RuneCountInString(receiver.value))

	return
}

// ______________________________________
func (receiver PrimitiveString) IsEmpty() (
	isEmpty bool,
) {
	isEmpty = receiver.Length() == 0

	return
}

// ______________________________________
func (receiver PrimitiveString) HasValue() (
	hasValue bool,
) {
	hasValue = !receiver.GetIsNil() && !receiver.IsEmpty()

	return
}

// ______________________________________
func (receiver PrimitiveString) Equal(
	value string,
) (
	ok bool,
) {
	ok = !receiver.GetIsNil() && receiver.value == value

	return
}

// ______________________________________
func (receiver PrimitiveString) Validation() (
	err error,
) {
	if receiver.GetIsNil() {
		err = nil

		return
	}

	if returnedErr := receiver.ValidationMax(); returnedErr != nil {
		err = returnedErr

		return
	}

	if returnedErr := receiver.ValidationMin(); returnedErr != nil {
		err = returnedErr

		return
	}

	err = receiver.ValidationSpell()

	return
}

// ValidationMax は最大文字列長のチェックを行います
// ______________________________________
func (receiver PrimitiveString) ValidationMax() (
	err error,
) {
	if receiver.GetIsNil() {
		err = nil

		return
	}

	if receiver.maxLength == nil {
		err = nil

		return
	}

	if receiver.Length() > *receiver.maxLength {
		err = receiver.newErrorString("max limitation")

		return
	}

	err = nil

	return
}

// ValidationMin は最小文字列長のチェックを行います
// ______________________________________
func (receiver PrimitiveString) ValidationMin() (
	err error,
) {
	if receiver.GetIsNil() {
		err = nil

		return
	}

	if receiver.minLength == nil {
		err = nil

		return
	}

	if receiver.Length() < *receiver.minLength {
		err = receiver.newErrorString("min limitation")

		return
	}

	err = nil

	return
}

// ValidationSpell は禁止文字列のチェックを行います
// ______________________________________
func (receiver PrimitiveString) ValidationSpell() (
	err error,
) {
	if len(receiver.spellList) == 0 {
		err = nil

		return
	}
	for _, spell := range receiver.spellList {
		if strings.Contains(receiver.value, spell) {
			err = receiver.newErrorString("detect target spell : " + spell)

			return
		}
	}

	err = nil

	return
}

// CheckNil は文字列ポインタのnilチェックを行い、適切な値を返します
// ______________________________________
func (receiver *PrimitiveString) CheckNil(
	value *string,
) (
	isNil bool,
) {
	isNil = true
	if value != nil {
		isNil = false
	}

	return
}
