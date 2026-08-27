package primitive_object

import (
	"fmt"
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
) PrimitiveStringOption {
	isNil := receiver.CheckNil(value)
	valueString := ""
	if !isNil {
		valueString = *value
	}

	return func(s *PrimitiveString) {
		s.value = valueString
		s.isNil = isNil
	}
}

// ______________________________________
// WithIsNil はnil状態を設定するオプションを返します
func (receiver *PrimitiveString) WithIsNil(
	isNil bool,
) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.isNil = isNil
	}
}

// ______________________________________
// WithMaxLength は最大文字列長を設定するオプションを返します
func (receiver *PrimitiveString) WithMaxLength(
	length *uint,
) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.maxLength = length
	}
}

// ______________________________________
// WithMinLength は最小文字列長を設定するオプションを返します
func (receiver *PrimitiveString) WithMinLength(
	length *uint,
) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.minLength = length
	}
}

// ______________________________________
// WithCheckSpell は禁止文字列リストを設定するオプションを返します
func (receiver *PrimitiveString) WithCheckSpell(
	spellList []string,
) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.spellList = spellList
	}
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
func (receiver PrimitiveString) GetIsNil() bool {
	return receiver.isNil
}

// ______________________________________
func (receiver PrimitiveString) newErrorString(
	errString string,
) error {
	return fmt.Errorf(
		"error: %s",
		errString,
	)
}

// ______________________________________
func (receiver PrimitiveString) GetValue() string {
	if receiver.GetIsNil() {
		return ""
	}
	return receiver.value
}

// ______________________________________
func (receiver PrimitiveString) Length() uint {
	if receiver.GetIsNil() {
		return 0
	}

	return uint(utf8.RuneCountInString(receiver.value))
}

// ______________________________________
func (receiver PrimitiveString) IsEmpty() bool {
	return receiver.Length() == 0
}

// ______________________________________
func (receiver PrimitiveString) HasValue() bool {
	return !receiver.GetIsNil() && !receiver.IsEmpty()
}

// ______________________________________
func (receiver PrimitiveString) Equal(value string) bool {
	return !receiver.GetIsNil() && receiver.value == value
}

// ______________________________________
func (receiver PrimitiveString) Validation() error {
	if receiver.GetIsNil() {
		return nil
	}

	if err := receiver.ValidationMax(); err != nil {
		return err
	}

	if err := receiver.ValidationMin(); err != nil {
		return err
	}

	return receiver.ValidationSpell()
}

// ValidationMax は最大文字列長のチェックを行います
// ______________________________________
func (receiver PrimitiveString) ValidationMax() error {
	if receiver.GetIsNil() {
		return nil
	}

	if receiver.maxLength == nil {
		return nil
	}

	if receiver.Length() > *receiver.maxLength {
		return receiver.newErrorString("max limitation")
	}

	return nil
}

// ValidationMin は最小文字列長のチェックを行います
// ______________________________________
func (receiver PrimitiveString) ValidationMin() error {
	if receiver.GetIsNil() {
		return nil
	}

	if receiver.minLength == nil {
		return nil
	}

	if receiver.Length() < *receiver.minLength {
		return receiver.newErrorString("min limitation")
	}

	return nil
}

// ValidationSpell は禁止文字列のチェックを行います
// ______________________________________
func (receiver PrimitiveString) ValidationSpell() error {
	if len(receiver.spellList) == 0 {
		return nil
	}
	for _, spell := range receiver.spellList {
		if strings.Contains(receiver.value, spell) {
			return receiver.newErrorString("detect target spell : " + spell)
		}
	}

	return nil
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
