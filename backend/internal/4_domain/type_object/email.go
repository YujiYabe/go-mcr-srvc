package type_object

import (
	"fmt"
	"regexp"

	primitiveObject "backend/internal/4_domain/primitive_object"
)

var (
	emailMaxLength uint = 30
	emailMinLength uint = 1
)

var emailCheckSpell = []string{}

const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

var emailRegexp = regexp.MustCompile(emailPattern)

type Email struct {
	content *primitiveObject.PrimitiveString
}

func NewEmail(
	value *string,
) (
	email Email,
	err error,
) {
	email = Email{}
	err = email.setValue(value)

	return
}

func (receiver *Email) setValue(
	value *string,
) error {
	primitiveString := &primitiveObject.PrimitiveString{}

	receiver.content = primitiveObject.NewPrimitiveString(
		primitiveString.WithValue(value),
		primitiveString.WithMaxLength(&emailMaxLength),
		primitiveString.WithMinLength(&emailMinLength),
		primitiveString.WithCheckSpell(emailCheckSpell),
	)
	if err := receiver.content.Validation(); err != nil {
		return err
	}

	return receiver.Validation()
}
func (receiver Email) GetValue() string {
	return receiver.content.GetValue()
}

func (receiver *Email) ErrorString(
	errString string,
) error {
	return fmt.Errorf("error: %s", errString)
}

func (receiver Email) GetIsNil() bool {
	return receiver.content.GetIsNil()
}

func (receiver Email) Validation() error {
	if receiver.GetIsNil() {
		return nil
	}

	if !emailRegexp.MatchString(receiver.GetValue()) {
		return fmt.Errorf("invalid email format: %s", receiver.GetValue())
	}
	return nil
}
