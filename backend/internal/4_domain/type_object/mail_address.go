package type_object

import (
	"fmt"
	"regexp"

	primitiveObject "backend/internal/4_domain/primitive_object"
)

var (
	mailAddressMaxLength uint = 30
	mailAddressMinLength uint = 1
)

var mailAddressCheckSpell = []string{}

type MailAddress struct {
	content *primitiveObject.PrimitiveString
}

func NewMailAddress(
	value *string,
) (
	mailAddress MailAddress,
	err error,
) {
	mailAddress = MailAddress{}
	err = mailAddress.SetValue(value)

	return
}

func (receiver *MailAddress) SetValue(
	value *string,
) error {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&mailAddressMaxLength),
		primitiveString.WithMinLength(&mailAddressMinLength),
		primitiveString.WithCheckSpell(mailAddressCheckSpell),
	)

	receiver.content.Validation()
	if receiver.content.GetError() != nil {
		return receiver.content.GetError()
	}

	return receiver.Validation()
}
func (receiver *MailAddress) GetValue() string {
	return receiver.content.GetValue()
}

func (receiver *MailAddress) ErrorString(
	errString string,
) error {
	return fmt.Errorf("error: %s", errString)
}

func (receiver *MailAddress) GetIsNil() bool {
	return receiver.content.GetIsNil()
}

func (receiver *MailAddress) Validation() error {
	if receiver.GetIsNil() {
		return nil
	}

	// メールアドレスの正規表現パターン
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	matched, err := regexp.MatchString(
		emailPattern,
		receiver.GetValue(),
	)
	if err != nil {
		return fmt.Errorf("failed to validate email format: %w", err)
	}

	if !matched {
		return fmt.Errorf("invalid email format: %s", receiver.GetValue())
	}
	return nil
}
