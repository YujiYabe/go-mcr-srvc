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
	err       error    // バリデーションエラーを格納
	value     string   // 実際の文字列値
	isNil     bool     // nil状態を示すフラグ
	maxLength int      // 最大文字列長 (-1は制限なし)
	minLength int      // 最小文字列長 (-1は制限なし)
	spellList []string // チェック対象の禁止文字列リスト
}

// ______________________________________
// NewPrimitiveString は指定されたオプションで新しいPrimitiveStringインスタンスを生成します
type PrimitiveStringOption func(*PrimitiveString)

// ______________________________________
// WithError はエラーを設定するオプションを返します
func (receiver *PrimitiveString) WithError(
	err error,
) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.err = err
	}
}

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
	length int,
) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.maxLength = length
	}
}

// ______________________________________
// WithMinLength は最小文字列長を設定するオプションを返します
func (receiver *PrimitiveString) WithMinLength(
	length int,
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
		err:       nil,
		value:     "",
		isNil:     false,
		maxLength: -1,
		minLength: -1,
		spellList: []string{},
	}

	// オプションを適用
	for _, option := range options {
		option(primitiveString)
	}

	return
}

// ______________________________________
func (receiver *PrimitiveString) SetIsNil(isNil bool) {
	receiver.isNil = isNil
}

// ______________________________________
func (receiver *PrimitiveString) GetIsNil() bool {
	return receiver.isNil
}

// ______________________________________
func (receiver *PrimitiveString) GetError() error {
	return receiver.err
}

// ______________________________________
func (receiver *PrimitiveString) SetErrorString(
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
func (receiver *PrimitiveString) SetError(
	err error,
) {
	receiver.err = err
}

// ______________________________________
func (receiver *PrimitiveString) GetValue() string {
	if receiver.GetIsNil() {
		return ""
	}
	return receiver.value
}

// ______________________________________
func (receiver *PrimitiveString) SetValue(
	value *string,
) {
	if value == nil {
		receiver.SetIsNil(true)
		return
	}
	receiver.SetIsNil(false)
	receiver.value = *value
}

// Validation は全てのバリデーションチェックを実行します
// ______________________________________
func (receiver *PrimitiveString) Validation() {

	if receiver.GetIsNil() {
		return
	}

	receiver.ValidationMax()
	if receiver.GetError() != nil {
		return
	}

	receiver.ValidationMin()
	if receiver.GetError() != nil {
		return
	}

	receiver.ValidationSpell()
	if receiver.GetError() != nil {
		return
	}

}

// ValidationMax は最大文字列長のチェックを行います
// ______________________________________
func (receiver *PrimitiveString) ValidationMax() {
	if receiver.maxLength < 0 {
		// receiver.SetError("max length no defined")
		return
	}

	if receiver.GetIsNil() {
		receiver.SetErrorString("is nil")
		return
	}

	if utf8.RuneCountInString(receiver.value) > receiver.maxLength {
		receiver.SetErrorString("max limitation")
		return
	}
}

// ValidationMin は最小文字列長のチェックを行います
// ______________________________________
func (receiver *PrimitiveString) ValidationMin() {
	if receiver.GetIsNil() {
		return
	}

	if receiver.minLength < 0 {
		return
	}

	if utf8.RuneCountInString(receiver.value) < receiver.minLength {
		receiver.SetErrorString("min limitation")
		return
	}
}

// ValidationSpell は禁止文字列のチェックを行います
// ______________________________________
func (receiver *PrimitiveString) ValidationSpell() {
	if len(receiver.spellList) == 0 {
		return
	}
	for _, spell := range receiver.spellList {
		if strings.Contains(receiver.value, spell) {
			receiver.SetErrorString("detect target spell : " + spell)
			return
		}
	}
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
