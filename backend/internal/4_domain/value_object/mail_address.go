package value_object

import (
	"backend/internal/4_domain/primitive_object"
	"fmt"
	"regexp"
)

const (
	mailAddressLengthMax = 30
	mailAddressLengthMin = 1
)

var mailAddressCheckSpell = []string{}

type MailAddress struct {
	err     error
	content *primitive_object.PrimitiveString
}

func NewMailAddress(
	value *string,
) (
	mailAddress MailAddress,
) {
	mailAddress = MailAddress{}
	mailAddress.SetValue(value)

	return
}

func (receiver *MailAddress) SetValue(
	value *string,
) {
	// 値の格納前にバリデーション。
	primitiveString := &primitive_object.PrimitiveString{}

	receiver.content = primitive_object.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(mailAddressLengthMax),
		primitiveString.WithMinLength(mailAddressLengthMin),
		primitiveString.WithCheckSpell(mailAddressCheckSpell),
	)

	// 文字列そのもののバリデーション
	receiver.Validation()
	if receiver.GetError() != nil {
		receiver.SetError(receiver.GetError())
		return
	}

	// メールアドレスのバリデーション
	receiver.Validation()

}

func (receiver *MailAddress) GetValue() string {
	return receiver.content.GetValue()
}

func (receiver *MailAddress) GetError() error {
	return receiver.content.GetError()
}

func (receiver *MailAddress) SetError(
	err error,
) {
	receiver.err = err
}

func (receiver *MailAddress) SetErrorString(
	errString string,
) {
	receiver.SetError(
		fmt.Errorf(
			"error: %s",
			errString,
		),
	)
}

func (receiver *MailAddress) GetIsNil() bool {
	return receiver.content.GetIsNil()
}

func (receiver MailAddress) Validation() {
	if receiver.GetIsNil() {
		return
	}

	// メールアドレスの正規表現パターン
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`

	matched, err := regexp.MatchString(
		emailPattern,
		receiver.GetValue(),
	)
	if err != nil {
		receiver.SetError(
			fmt.Errorf(
				"failed to validate email format: %w", err,
			),
		)
		return
	}

	if !matched {
		receiver.SetError(
			fmt.Errorf(
				"invalid email format: %s", receiver.GetValue(),
			),
		)
		return
	}
}
