package primitive_object

import (
	"fmt"
	"log"
	"strings"
	"unicode/utf8"
)

type ContextKey string

// --------------------------------------
// PrimitiveString は文字列値に対してバリデーション機能を提供する構造体です。
// nil チェック、長さ制限、禁止文字列のチェックなどの機能を備えています。
type PrimitiveString struct {
	err       error    // バリデーションエラーを格納
	value     string   // 実際の文字列値
	isNil     bool     // nil状態を示すフラグ
	MaxLength int      // 最大文字列長 (-1は制限なし)
	MinLength int      // 最小文字列長 (-1は制限なし)
	SpellList []string // チェック対象の禁止文字列リスト
}

// --------------------------------------
// NewPrimitiveString は指定されたオプションで新しいPrimitiveStringインスタンスを生成します
type PrimitiveStringOption func(*PrimitiveString)

// --------------------------------------
// WithError はエラーを設定するオプションを返します
func (receiver *PrimitiveString) WithError(
	err error,
) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.err = err
	}
}

// --------------------------------------
// WithValue は文字列値を設定するオプションを返します
func (receiver *PrimitiveString) WithValue(
	value *string,
) PrimitiveStringOption {
	receiver.SetIsNil(true)
	var resValue string
	if value != nil {
		receiver.SetIsNil(false)
		resValue = *value
	}
	return func(s *PrimitiveString) {
		s.value = resValue
	}
}

// --------------------------------------
// WithIsNil はnil状態を設定するオプションを返します
func (receiver *PrimitiveString) WithIsNil(
	isNil bool,
) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.isNil = isNil
	}
}

// --------------------------------------
// WithMaxLength は最大文字列長を設定するオプションを返します
func (receiver *PrimitiveString) WithMaxLength(
	length int,
) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.MaxLength = length
	}
}

// --------------------------------------
// WithMinLength は最小文字列長を設定するオプションを返します
func (receiver *PrimitiveString) WithMinLength(
	length int,
) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.MinLength = length
	}
}

// --------------------------------------
// WithCheckSpell は禁止文字列リストを設定するオプションを返します
func (receiver *PrimitiveString) WithCheckSpell(
	spellList []string,
) PrimitiveStringOption {
	return func(s *PrimitiveString) {
		s.SpellList = spellList
	}
}

// --------------------------------------
func NewPrimitiveString(
	options ...PrimitiveStringOption,
) (
	primitiveString *PrimitiveString,
) {

	var defaultValue string
	var defaultIsNil bool

	var defaultMaxLength int = -1
	var defaultMinLength int = -1

	// デフォルト値を設定
	primitiveString = &PrimitiveString{
		err:       nil,
		value:     defaultValue,
		isNil:     defaultIsNil,
		MaxLength: defaultMaxLength,
		MinLength: defaultMinLength,
		SpellList: []string{},
	}

	// オプションを適用
	for _, option := range options {
		option(primitiveString)
	}

	return
}

// --------------------------------------
func (receiver *PrimitiveString) SetIsNil(isNil bool) {
	receiver.isNil = isNil
}

// --------------------------------------
func (receiver *PrimitiveString) GetIsNil() bool {
	return receiver.isNil
}

// --------------------------------------
func (receiver *PrimitiveString) GetError() error {
	return receiver.err
}

// --------------------------------------
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

// --------------------------------------
func (receiver *PrimitiveString) SetError(
	err error,
) {
	receiver.err = err
}

// --------------------------------------
func (receiver *PrimitiveString) GetValue() string {
	if receiver.isNil {
		receiver.SetErrorString("is nil")
		return ""
	}
	return receiver.value
}

// --------------------------------------
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
// --------------------------------------
func (receiver *PrimitiveString) Validation() {

	if receiver.isNil {
		return
	}

	receiver.ValidationMax()
	if receiver.err != nil {
		return
	}

	receiver.ValidationMin()
	if receiver.err != nil {
		return
	}

	receiver.ValidationSpell()
	if receiver.err != nil {
		return
	}

}

// ValidationMax は最大文字列長のチェックを行います
// --------------------------------------
func (receiver *PrimitiveString) ValidationMax() {
	if receiver.MaxLength < 0 {
		// receiver.SetError("max length no defined")
		return
	}

	if receiver.isNil {
		receiver.SetErrorString("is nil")
		return
	}

	if utf8.RuneCountInString(receiver.value) > receiver.MaxLength {
		receiver.SetErrorString("max limitation")
		return
	}
}

// ValidationMin は最小文字列長のチェックを行います
// --------------------------------------
func (receiver *PrimitiveString) ValidationMin() {
	log.Println("==receiver.GetIsNil() == == == == == == == == == ")
	log.Printf("%#v\n", receiver.GetIsNil())
	log.Println("== == == == == == == == == == ")

	if !receiver.GetIsNil() {
		return
	}

	if receiver.MinLength < 0 {
		return
	}

	log.Println("== == == == == == == == == == ")
	log.Printf("%#v\n", utf8.RuneCountInString(receiver.value))
	log.Printf("%#v\n", receiver.MinLength)
	log.Println("== == == == == == == == == == ")

	if utf8.RuneCountInString(receiver.value) < receiver.MinLength {

		receiver.SetErrorString("min limitation")
		return
	}
}

// ValidationSpell は禁止文字列のチェックを行います
// --------------------------------------
func (receiver *PrimitiveString) ValidationSpell() {
	if len(receiver.SpellList) == 0 {
		return
	}
	for _, spell := range receiver.SpellList {
		if strings.Contains(receiver.value, spell) {
			receiver.SetErrorString("detect target spell : " + spell)
			return
		}
	}
}

// CheckNil は文字列ポインタのnilチェックを行い、適切な値を返します
// --------------------------------------
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
